package routers

import (
	"quiz-admin/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.MainController{})
	beego.Include(&controllers.AccountController{})
	beego.Include(&controllers.AppController{})
}
