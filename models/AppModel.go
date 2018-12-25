package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func AppGetAll(offset int) DBResponse {
	var maps []orm.Params
	eMsg := ""
	var res DBResponse

	o := orm.NewOrm()

	query := `SELECT *
	 FROM tbl_app 
	 order by updated_at desc`
	query += " OFFSET ? "
	if offset > 0 {
		query += " FETCH NEXT 10 ROWS only;"
	} else {
		query += " limit 25;"
	}
	num, err := o.Raw(query, offset).Values(&maps)
	if err != nil {
		beego.Error(err)
		eMsg = err.Error()
	}
	res = DBResponse{num, eMsg, maps}
	return res
}

func AppAdd(app_name, gateway_app_id, gateway_name, gateway_secret, country, currency, language, privacy_policy, term_condition, otp_attempt, app_icon, status, total_question string) DBResponse {
	var maps []orm.Params
	eMsg := ""
	o := orm.NewOrm()
	num, err := o.Raw("INSERT INTO tbl_app "+
		"(app_name, gateway_app_id, gateway_name, gateway_secret, country, currency, language, privacy_policy, term_condition, otp_attempt, app_icon, status, total_question, created_by, created_at, updated_by, updated_at) "+
		"VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,'admin',NOW(),'admin',NOW());", app_name, gateway_app_id, gateway_name, gateway_secret, country, currency, language, privacy_policy, term_condition, otp_attempt, app_icon, status, total_question).Values(&maps)
	//beego.Info("Total Rows found : ", num)
	if err != nil {
		beego.Error(err)
		eMsg = err.Error()
	}
	for k, v := range maps {
		beego.Info("k:", k, "v:", v)
	}
	return DBResponse{num, eMsg, maps}
}

func AppGetById(app_id string) DBResponse {
	var maps []orm.Params
	eMsg := ""
	var res DBResponse
	o := orm.NewOrm()
	num, err := o.Raw("select * from tbl_app where app_id = ? limit 1;", app_id).Values(&maps)
	if err != nil {
		beego.Error(err)
		eMsg = err.Error()
	}
	res = DBResponse{num, eMsg, maps}
	return res
}

func AppEdit(app_id, app_name, gateway_app_id, gateway_name, gateway_secret, country, currency, language, privacy_policy, term_condition, otp_attempt, app_icon, status, total_question string) DBResponse {
	eMsg := ""
	var num int64 = 0
	o := orm.NewOrm()
	query := "UPDATE tbl_app SET "
	if app_icon != "" {
		query += "app_icon = '" + app_icon + "',"
	}
	query += "app_name = ?, gateway_app_id = ?, gateway_name = ?, gateway_secret = ?, country = ?,currency = ?, language = ?,privacy_policy = ?, term_condition = ?,otp_attempt = ?, status = ?, total_question = ?, created_by = 'admin' , updated_by = 'admin', updated_at= now() "

	query += " where app_id = ?;"

	res, err := o.Raw(query, app_name, gateway_app_id, gateway_name, gateway_secret, country, currency, language, privacy_policy, term_condition, otp_attempt, status, total_question, app_id).Exec()
	if err != nil {
		beego.Error(err)
		eMsg = err.Error()
	} else {
		num, _ := res.RowsAffected()
		beego.Info("Total Rows affected : ", num)
	}
	return DBResponse{num, eMsg, nil}
}

func AppIdGet() DBResponse {
	var maps []orm.Params
	eMsg := ""
	var res DBResponse
	o := orm.NewOrm()

	query := `SELECT app_id, app_name
	FROM tbl_app;`

	num, err := o.Raw(query).Values(&maps)
	if err != nil {
		beego.Error(err)
		eMsg = err.Error()
	}

	res = DBResponse{num, eMsg, maps}
	return res
}
