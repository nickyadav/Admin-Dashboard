package models

import (
	"encoding/json"
	"time"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
)

var bm cache.Cache
var err error
var StoreString string

//Constant for Cache
var CACHE_KEY_COUNTRY_SINGLE string = "country_"
var CACHE_KEY_COUNTRY_ALL string = "countries"

var CACHE_SHORT_INTERVAL = 1 * time.Minute
var CACHE_MEDIUM_INTERVAL = 5 * time.Minute
var CACHE_LONG_INTERVAL = 15 * time.Minute
var CACHE_LONGEST_INTERVAL = 1 * time.Hour

var LIMIT_MINIMUM = "10"
var LIMIT_MEDIUM = "25"
var LIMIT_MAXIMUM = "100"

var PAGINATION_MINIMUM = 10
var PAGINATION_MEDIUM = 50
var PAGINATION_MAXIMUM int = 200

type DBResponse struct {
	Rows         int64
	ErrorMessage string
	Data         []orm.Params
}

type Pagination struct {
	Previous int
	Current  int
	Next     int
}

type GlobalSession struct {
	Id        string
	Username  string
	Fullname  string
	RoleID    string
	RoleTitle string
	RoleName  string
	LoginAt   time.Time
}

func init() {
	bm, err = cache.NewCache("memory", `{"interval":3600}`)
	//orm.RegisterModelWithPrefix("tbl_", new(RuleConfig))
}

func GetGlobalSession(ses string) GlobalSession {
	res := GlobalSession{}
	json.Unmarshal([]byte(ses), &res)
	return res
}
