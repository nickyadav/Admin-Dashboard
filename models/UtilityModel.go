package models

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func EncryptPassword(str string) string {
	password := []byte(str)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, 15)
	if err != nil {
		beego.Error(err)
	}
	return string(hashedPassword)
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

// generates a random string of fixed size
func RandomString(size int) string {
	var alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890+_-"
	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		buf[i] = alpha[rand.Intn(len(alpha))]
	}
	return string(buf)
}

func ConvertToTime(dt string) (time.Time, error) {
	//layout := "2006-01-02 15:04:05 -0700 MST"
	//t, err := time.Parse(layout, "2014-11-17 23:02:03 +0000 UTC")
	//fmt.Println(t, err)
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, dt)
	beego.Info(t, err)
	return t, err
}

func ConvertToInt(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	beego.Info(i, err)
	return i, err
}

func ConvertToViewDate(dt string) string {
	t, _ := time.Parse(time.RFC3339, dt)
	dateTime := t.Format("2006-01-02 15:04:05 AM")
	return dateTime
}

func GetDateRange() {
	startDate, _ := time.Parse("2006-01-02 15:04:05", "2018-08-08 15:48:11")
	endDate, _ := time.Parse("2006-01-02 15:04:05", "2018-09-08 15:48:15")
	beego.Info(startDate, " - ", endDate)
	duration := endDate.Sub(startDate)
	days := int(duration.Hours()/24) + 1
	beego.Info(days)
	for i := 0; i < days; {
		if (i + 7) < days {
			beego.Info("S : ", startDate.AddDate(0, 0, i).Format("2006-01-02"), " - E : ", startDate.AddDate(0, 0, i+6).Format("2006-01-02"))
		} else {
			beego.Info("S : ", startDate.AddDate(0, 0, i).Format("2006-01-02"), " - E : ", endDate.Format("2006-01-02"))
		}
		i = i + 7
	}
}

func GetMD5Hash(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GetSHA256Hash(input string) string {
	sha_256 := sha256.New()
	sha_256.Write([]byte(input))
	return hex.EncodeToString(sha_256.Sum(nil))
}
