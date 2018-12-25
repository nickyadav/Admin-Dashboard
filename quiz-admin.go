package main

import (
	"quiz-admin/controllers"

	_ "quiz-admin/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	beego.BConfig.Log.AccessLogs = true
	orm.Debug = true
	orm.RegisterDriver("postgres", orm.DRPostgres)
	dataSource := "user=" + beego.AppConfig.String("db_user") +
		" password=" + beego.AppConfig.String("db_pass") +
		" host=" + beego.AppConfig.String("db_host") +
		" port=" + beego.AppConfig.String("db_port") +
		" dbname=" + beego.AppConfig.String("db_name") +
		" sslmode=" + beego.AppConfig.String("db_sslmode") +
		" connect_timeout=" + beego.AppConfig.String("db_connect_timeout")
	//beego.Info(dataSource)
	orm.RegisterDataBase("default", "postgres", dataSource)
	orm.RunCommand()
}

func main() {
	beego.SetStaticPath("/bower_components", "bower_components")
	logs.SetLogger(logs.AdapterFile, `{"filename":"logs/log.txt"}`)
	beego.InsertFilter("/*", beego.BeforeRouter, controllers.FilterUser)
	// beego.AddFuncMap("ChangeDateFormat", controllers.ChangeDateFormat)
	//beego.ErrorController(&controllers.ErrorController{})
	// models.S3Upload("temp/abc.png", "")
	beego.Run()
}
