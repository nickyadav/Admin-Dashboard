package clients

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

func CompletePathCreate(ImgName, gatewayName, controllerName string) (string, string) {
	var completeImgPath string
	var filePath string
	t := time.Now()
	timestr := strconv.FormatInt(time.Now().Unix(), 10)
	beego.Info("Time ", timestr)
	year := strconv.Itoa(t.Year())
	var mont int = int(t.Month())
	month := fmt.Sprintf("%02d", mont)
	date := strconv.Itoa(t.Day())

	if ImgName == "" {
		beego.Info("No reward image has been provided.")
	} else {
		filePath += "temp/" + ImgName
		path := year + "/" + month + "/" + date + "/" + timestr + ImgName
		completeImgPath += gatewayName + "/" + controllerName + "/" + path
	}
	return completeImgPath, filePath
}
