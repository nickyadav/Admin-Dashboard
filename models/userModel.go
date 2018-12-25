package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func UserGetAll(offset int) DBResponse {
	var maps []orm.Params
	eMsg := ""
	var res DBResponse
	o := orm.NewOrm()

	query := `SELECT user_id, username, email, user_fullname, status, created_by,  updated_by, role_id 
	FROM tbl_user  
	order by username `
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

func UserAdd(username, passwordHash, email, userFullName, status string) DBResponse {
	var maps []orm.Params
	eMsg := ""
	o := orm.NewOrm()
	num, err := o.Raw("INSERT into tbl_user "+
		"(username, password_hash, email, user_fullname, status, role_id, created_by, created_at, updated_by, updated_at) "+
		"VALUES(?,?,?,?,?,Null,'admin',NOW(),'admin',NOW()) RETURNING user_id;", username, passwordHash, email, userFullName, status).Values(&maps)
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

func UserGetById(id string) DBResponse {
	var maps []orm.Params
	eMsg := ""
	o := orm.NewOrm()
	num, err := o.Raw("select user_id, username, email, user_fullname, status, role_id, created_by, updated_by from tbl_user where user_id = ? limit 1;", id).Values(&maps)
	//beego.Info("Total Rows found : ", num)
	if err != nil {
		beego.Error(err)
		eMsg = err.Error()
	}
	//for k, v := range maps {
	//   beego.Info("k:", k, "v:", v)
	//}
	return DBResponse{num, eMsg, maps}
}

func UserGetByUsername(username string) DBResponse {
	var maps []orm.Params
	eMsg := ""
	o := orm.NewOrm()

	num, err := o.Raw(`select user_id, username, user_fullname, u.role_id, r.role_title, r.role_name, email, password_hash, status, u.created_at, u.created_by, u.updated_at, u.updated_by 
		from tbl_user u 
		left join tbl_role r on r.role_id = u.role_id 
		where status = 'ACTIVE' AND  u.username= ?  limit 1;`, username).Values(&maps)
	//beego.Info("Total Rows found : ", num)
	if err != nil {
		beego.Error(err)
		eMsg = err.Error()
	}
	return DBResponse{num, eMsg, maps}
}

func UserEdit(id, username, email, password, fullname, status string) DBResponse {
	eMsg := ""
	var num int64 = 0
	o := orm.NewOrm()
	res, err := o.Raw("UPDATE tbl_user "+
		"SET username = ?, email = ?, password_hash = ?, user_fullname = ?, status = ?, updated_by = 'admin', updated_at = now()"+
		" where user_id = ?;",
		username, email, password, fullname, status, id).Exec()
	if err != nil {
		beego.Error(err)
		eMsg = err.Error()
	} else {
		num, _ := res.RowsAffected()
		beego.Info("Total Rows affected : ", num)
	}
	return DBResponse{num, eMsg, nil}
}

func UserTokenCreate(username, resetToken string, resetTime time.Time) DBResponse {
	eMsg := ""
	var num int64 = 0
	o := orm.NewOrm()
	res, err := o.Raw("UPDATE tbl_user "+
		"SET password_reset_token = ?, reset_token_expire = ? "+
		" where username = ?;",
		resetToken, resetTime, username).Exec()
	if err != nil {
		beego.Error(err)
		eMsg = err.Error()
	} else {
		num, _ := res.RowsAffected()
		beego.Info("Total Rows affected : ", num)
	}
	return DBResponse{num, eMsg, nil}
}

func UserTokenCheck(resetToken string) DBResponse {
	var maps []orm.Params
	eMsg := ""
	o := orm.NewOrm()

	num, err := o.Raw(`select reset_token_expire
		from tbl_user  
		where password_reset_token = ? limit 1;`, resetToken).Values(&maps)
	//beego.Info("Total Rows found : ", num)
	if err != nil {
		beego.Error(err)
		eMsg = err.Error()
	}
	return DBResponse{num, eMsg, maps}
}

func UserPasswordUpdate(passwordResetToken, resetPassword string) DBResponse {
	eMsg := ""
	var num int64 = 0
	o := orm.NewOrm()
	res, err := o.Raw("UPDATE tbl_user "+
		"SET password_hash = ? "+
		" where password_reset_token = ?;",
		resetPassword, passwordResetToken).Exec()
	if err != nil {
		beego.Error(err)
		eMsg = err.Error()
	} else {
		num, _ := res.RowsAffected()
		beego.Info("Total Rows affected : ", num)
	}
	return DBResponse{num, eMsg, nil}
}
