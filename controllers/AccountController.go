package controllers

import (
	"strconv"
	"time"

	"quiz-admin/models"

	"github.com/astaxie/beego"
)

type AccountController struct {
	baseRouter
}

func (c *AccountController) NestPrepare() {
}

// @router /account/login [get]
func (c *AccountController) LoginGet() {
	c.TplName = "account/login.html"
}

// @router /account/login [post]
func (c *AccountController) LoginPost() {
	flash := beego.NewFlash()
	//Done: Step 1 - Get User
	username := c.GetString("username")
	password := c.GetString("password")
	r := c.GetString("r")
	beego.Info("Redirect to : " + r)
	//Done: Step 2 - Check In DB
	dbRes := models.UserGetByUsername(username)
	if dbRes.Rows > 0 {
		//Done: Step 3 - Validate Password
		if models.ComparePassword(dbRes.Data[0]["password_hash"].(string), password) {
			//Done: Step 4 - If Validated Store the important value to Session & redirect to dashboard
			var gs models.GlobalSession
			gs.Id = dbRes.Data[0]["user_id"].(string)
			gs.Username = dbRes.Data[0]["username"].(string)
			gs.Fullname = dbRes.Data[0]["user_fullname"].(string)

			c.SetSession("GlobalSession", models.StructToJSONStr(gs))
			if r != "" {
				c.Redirect(r, 302)
			} else {
				c.Redirect(c.URLFor("MainController.Index"), 302)
			}
		} else {
			beego.Error("Login Fail for : '" + username + "' Due to Password mismatch!")
		}
	} else {
		beego.Error("Login Fail for : '" + username + "' as no user found for the provided username!")
	}
	//Done: Step 5 - If validated failed show the login page with error message.
	flash.Error("Invalid Credential")
	flash.Store(&c.Controller)
	c.TplName = "account/login.html"
}

// @router /account/logout [get]
func (c *AccountController) Logout() {
	c.DelSession("GlobalSession")
	c.Redirect(c.URLFor("MainController.Index"), 302)
}

// @router /account/reset-password/:resetToken [get]
func (c *AccountController) ResetPasswordGet() {
	flash := beego.NewFlash()
	resetToken := c.Ctx.Input.Param(":resetToken")
	token := models.UserTokenCheck(resetToken)
	currentTime := time.Now()
	tokenExpireTime, _ := time.Parse(time.RFC3339, token.Data[0]["reset_token_expire"].(string))
	if tokenExpireTime.After(currentTime) { //check if currentTime is less then expireTime from db
		c.TplName = "account/reset-password.html"
	} else if tokenExpireTime.Before(currentTime) {
		flash.Error("Reset token expired redirecting to login page!!")
		flash.Store(&c.Controller)
		c.TplName = "account/login.html"
	}
}

// @router /account/reset-password/:resetToken [post]
func (c *AccountController) ResetPasswordPost() {
	flash := beego.NewFlash()
	resetToken := c.Ctx.Input.Param(":resetToken")
	resetPassword := c.GetString("resetPassword")
	resetConfirmPassword := c.GetString("confirmPassword")
	errFlag := false
	if resetPassword != resetConfirmPassword {
		flash.Error("Password didn't matched! Try again.")
		flash.Store(&c.Controller)
		c.TplName = "account/reset-password.html"
		errFlag = true
	}
	if !errFlag {
		models.UserPasswordUpdate(resetToken, models.EncryptPassword(resetPassword))
		flash.Store(&c.Controller)
		c.TplName = "account/login.html"
	}
}

// @router /profile [get]
func (c *AccountController) ProfileGet() {
	Username := c.GSession.Username
	user := models.UserGetByUsername(Username)
	c.Data["data"] = user.Data[0]
	c.TplName = "profile.html"
}

// @router /profile [post]
func (c *AccountController) ProfilePost() {
	flash := beego.NewFlash()
	username := c.GSession.Username
	oldPassword := c.GetString("oldPassword")
	newPassword := c.GetString("newPassword")
	confirmPassword := c.GetString("confirmPassword")
	email := c.GetString("email")
	fullname := c.GetString("user_fullname")
	errFlag := false
	if newPassword != confirmPassword {
		flash.Error("Password didn't matched! Try again.")
		flash.Store(&c.Controller)
		errFlag = true
	}

	if !errFlag {
		dbRes := models.UserGetByUsername(username)
		if dbRes.Rows > 0 {
			if models.ComparePassword(dbRes.Data[0]["password_hash"].(string), oldPassword) {
				models.UserEdit(c.GetString("user_id"), username, email, models.EncryptPassword(newPassword), fullname, "ACTIVE")

			} else {
				flash.Error("Error: Due to Password mismatch!")
			}
		}
	}
	flash.Store(&c.Controller)
	c.TplName = "dashboard.html"
}

// @router /account/ [get]
func (c *AccountController) Index() {
	page := c.GetString("page")
	pageNumb, _ := strconv.Atoi(page) //convert string to int
	pagination := models.Pagination{pageNumb - 1, pageNumb, pageNumb + 1}
	offset := pageNumb * 10
	result := models.UserGetAll(offset)
	c.Data["data"] = result
	c.Data["page"] = pagination
	c.TplName = "account/list.html"
}

// @router /account/new [get]
func (c *AccountController) NewGet() {
	c.TplName = "account/new.html"
}

// @router /account/new [post]
func (c *AccountController) NewPost() {
	flash := beego.NewFlash()
	errFlag := false
	m := make(map[string]string)
	m["username"] = c.GetString("username")
	m["user_fullname"] = c.GetString("user_fullname")
	m["emailid"] = c.GetString("emailid")
	m["password"] = c.GetString("password")
	m["confirmpassword"] = c.GetString("confirmpassword")
	m["status"] = c.GetString("status")
	c.Data["data"] = m
	if m["password"] != m["confirmpassword"] {
		flash.Error("Password didn't matched! Try again.")
		errFlag = true
	}

	if !errFlag {
		result := models.UserAdd(m["username"], models.EncryptPassword(m["password"]), m["emailid"], m["user_fullname"], m["status"])
		if result.ErrorMessage != "" {
			flash.Error(result.ErrorMessage)
		} else {
			flash.Notice("User Successfully added!")
			flash.Store(&c.Controller)
			c.Redirect(c.URLFor("AccountController.Index"), 302)
		}
	}

	flash.Store(&c.Controller)
	c.TplName = "account/new.html"
}

// @router /account/:id [get]
func (c *AccountController) EditGet() {
	beego.Info(c.Ctx.Input.Param(":id"))
	result := models.UserGetById(c.Ctx.Input.Param(":id"))
	if result.Rows > 0 {
		c.Data["data"] = result.Data[0]
		c.Data["title"] = result.Data[0]["user_fullname"]
	}
	if result.Rows == 0 {
		c.Abort("404")
	}
	c.TplName = "account/edit.html"
}

// @router /account/:id [post]
func (c *AccountController) EditPost() {
	flash := beego.NewFlash()
	errFlag := false
	m := make(map[string]string)
	m["username"] = c.GetString("username")
	m["user_fullname"] = c.GetString("user_fullname")
	m["emailid"] = c.GetString("emailid")
	m["password"] = c.GetString("password")
	m["confirmpassword"] = c.GetString("confirmpassword")
	m["status"] = c.GetString("status")
	c.Data["data"] = m
	if m["password"] != m["confirmpassword"] {
		flash.Error("Password didn't matched! Try again.")
		errFlag = true
	}

	if !errFlag {
		result := models.UserEdit(c.Ctx.Input.Param(":id"), m["username"], m["emailid"], models.EncryptPassword(m["password"]), m["user_fullname"], m["status"])
		if result.ErrorMessage == "" {
			flash.Notice("User Updated!")
		} else {
			flash.Error(result.ErrorMessage)
		}
	}
	flash.Store(&c.Controller)
	c.Redirect(c.URLFor("AccountController.View", ":id", c.Ctx.Input.Param(":id")), 302)

}

// @router /account/view/:id [get]
func (c *AccountController) View() {
	beego.Info("Getting Information for UserId : " + c.Ctx.Input.Param(":id"))
	result := models.UserGetById(c.Ctx.Input.Param(":id"))
	if result.Rows > 0 {
		c.Data["data"] = result.Data[0]
		c.Data["title"] = result.Data[0]["user_fullname"]
	}
	if result.Rows == 0 {
		c.Abort("404")
	}
	c.TplName = "account/view.html"
}
