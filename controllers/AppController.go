package controllers

import (
	"strconv"

	clients "quiz-admin/client"
	"quiz-admin/models"

	"github.com/astaxie/beego"
)

type AppController struct {
	baseRouter
}

// @router /app/ [get]
func (c *AppController) Index() {
	page := c.GetString("page")
	pageNumb, _ := strconv.Atoi(page) //convert string to int
	pagination := models.Pagination{pageNumb - 1, pageNumb, pageNumb + 1}
	offset := pageNumb * 10
	result := models.AppGetAll(offset)
	c.Data["data"] = result
	c.Data["page"] = pagination
	c.TplName = "app/list.html"
}

// @router /app/new [get]
func (c *AppController) NewGet() {
	c.TplName = "app/new.html"
}

// @router /app/new [post]
func (c *AppController) NewPost() {
	flash := beego.NewFlash()
	m := make(map[string]string)
	m["app_name"] = c.GetString("app_name")
	m["gateway_app_id"] = c.GetString("gateway_app_id")
	m["gateway_name"] = c.GetString("gateway_name")
	m["gateway_secret"] = c.GetString("gateway_secret")
	m["country"] = c.GetString("country")
	m["currency"] = c.GetString("currency")
	m["language"] = c.GetString("language")
	m["privacy_policy"] = c.GetString("privacy_policy")
	m["term_condition"] = c.GetString("term_condition")
	m["otp_attempt"] = c.GetString("otp_attempt")
	m["app_icon"] = c.GetString("app_icon")
	m["status"] = c.GetString("status")
	m["total_question"] = c.GetString("total_question")

	var completePath, filePath string
	_, imgName, _ := c.GetFile("app_icon")
	controllerName := "appIcon"
	if imgName == nil {
		beego.Info("No Image file has been provided.")
	} else {
		completePath, filePath = clients.CompletePathCreate(imgName.Filename, c.GetString("gateway_name"), controllerName)
	}
	result := models.AppAdd(m["app_name"], m["gateway_app_id"], m["gateway_name"], m["gateway_secret"], m["country"], m["currency"], m["language"], m["privacy_policy"], m["term_condition"], m["otp_attempt"], completePath, m["status"], m["total_question"])
	if result.ErrorMessage != "" {
		flash.Error(result.ErrorMessage)
	} else {
		if imgName == nil {
			beego.Info("No Image file has been provided.")
		} else {
			c.SaveToFile("app_icon", filePath)
		}
		flash.Notice("App Successfully added!")
		flash.Store(&c.Controller)
		c.Redirect(c.URLFor("AppController.Index"), 302)
	}

	flash.Store(&c.Controller)
	c.Data["data"] = m
	c.TplName = "app/new.html"
}

// @router /app/:id [get]
func (c *AppController) EditGet() {
	beego.Info(c.Ctx.Input.Param(":id"))
	result := models.AppGetById(c.Ctx.Input.Param(":id"))
	if result.Rows > 0 {
		c.Data["data"] = result.Data[0]
		if len(result.Data[0]["app_icon"].(string)) > 0 {
			url := result.Data[0]["app_icon"].(string)
			c.Data["url"] = url
		}
	}
	c.TplName = "app/edit.html"
}

// @router /app/:id [post]
func (c *AppController) EditPost() {
	flash := beego.NewFlash()
	m := make(map[string]string)
	m["app_name"] = c.GetString("app_name")
	m["gateway_app_id"] = c.GetString("gateway_app_id")
	m["gateway_name"] = c.GetString("gateway_name")
	m["gateway_secret"] = c.GetString("gateway_secret")
	m["country"] = c.GetString("country")
	m["currency"] = c.GetString("currency")
	m["language"] = c.GetString("language")
	m["privacy_policy"] = c.GetString("privacy_policy")
	m["term_condition"] = c.GetString("term_condition")
	m["otp_attempt"] = c.GetString("otp_attempt")
	m["app_icon"] = c.GetString("app_icon")
	m["status"] = c.GetString("status")
	m["total_question"] = c.GetString("total_question")

	var completePath, filePath string
	_, imgName, _ := c.GetFile("app_icon")
	controllerName := "appIcon"
	if imgName == nil {
		beego.Info("No Image file has been provided.")
	} else {
		completePath, filePath = clients.CompletePathCreate(imgName.Filename, c.GetString("gateway_name"), controllerName)
	}
	result := models.AppEdit(c.Ctx.Input.Param(":id"), m["app_name"], m["gateway_app_id"], m["gateway_name"], m["gateway_secret"], m["country"], m["currency"], m["language"], m["privacy_policy"], m["term_condition"], m["otp_attempt"], completePath, m["status"], m["total_question"])
	if result.ErrorMessage == "" {
		if imgName == nil {
			beego.Info("No Image file has been provided.")
		} else {
			c.SaveToFile("app_icon", filePath)
		}
		flash.Notice("App Updated!")
	} else {
		flash.Error(result.ErrorMessage)
	}

	flash.Store(&c.Controller)
	c.Redirect(c.URLFor("AppController.View", ":id", c.Ctx.Input.Param(":id")), 302)

}

// @router /app/view/:id [get]
func (c *AppController) View() {
	beego.Info("Getting Information for UserId : " + c.Ctx.Input.Param(":id"))
	result := models.AppGetById(c.Ctx.Input.Param(":id"))
	if result.Rows > 0 {
		c.Data["data"] = result.Data[0]
		c.Data["title"] = result.Data[0]["user_fullname"]
	}
	if result.Rows == 0 {
		c.Abort("404")
	}
	c.TplName = "app/view.html"
}
