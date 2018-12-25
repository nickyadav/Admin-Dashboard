package controllers

import (
	"html/template"
	"time"

	"quiz-admin/models"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

type NestPreparer interface {
	NestPrepare()
}

// baseRouter implements global settings for all other routers.
type baseRouter struct {
	beego.Controller
	i18n.Locale
	GSession models.GlobalSession
	isLogin  bool
}

// Prepare implements Prepare method for baseRouter.
func (c *baseRouter) Prepare() {
	beego.ReadFromRequest(&c.Controller)
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	_, ok := c.GetSession("GlobalSession").(string)
	if ok {
		// page start time
		c.Data["PageStartTime"] = time.Now()

		c.GSession = models.GetGlobalSession(c.GetSession("GlobalSession").(string))
		//beego.Info(c.GSession)
		c.Data["gSes"] = c.GSession
		c.Data["USER_NAME"] = c.GSession.Fullname
		c.Data["ROLE_ID"] = c.GSession.RoleID
		c.Data["ROLE_TITLE"] = c.GSession.RoleTitle
		c.Data["ROLE_NAME"] = c.GSession.RoleName
		//c.Data["SERVICE_PROVIDER_ID"] = c.GSession.ServiceProviderId
	}
	if app, ok := c.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}
