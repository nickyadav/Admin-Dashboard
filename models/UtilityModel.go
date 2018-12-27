package models

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func ComparePassword(strHashedPassword string, strPassword string) bool {
	flag := false
	hashedPassword := []byte(strHashedPassword)
	password := []byte(strPassword)
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err == nil {
		flag = true
	} else {
		beego.Error(err)
	}
	return flag
}

func EncryptPassword(str string) string {
	password := []byte(str)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, 15)
	if err != nil {
		beego.Error(err)
	}
	return string(hashedPassword)
}

func StructToJSONStr(value interface{}) string {
	data := &value
	js, err := json.Marshal(data)
	if err != nil {
		beego.Error(err)
		return ""
	}
	//beego.Info(string(js))
	return string(js)
}
