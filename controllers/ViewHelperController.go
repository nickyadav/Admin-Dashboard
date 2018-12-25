package controllers

import "github.com/astaxie/beego"

func GetbaseUrl(path string) string {
	url := beego.AppConfig.String("STATIC_FILE_URL")
	return url
}
