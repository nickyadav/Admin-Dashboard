package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func QuestionGetAll(pageNum int) DBResponse {
	var maps []orm.Params
	eMsg := ""
	var res DBResponse
	offset := pageNum * PAGINATION_MEDIUM
	o := orm.NewOrm()

	query := `SELECT q.*, c.category_name
	FROM tbl_question q
	join tbl_category c on c.category_id = q.category_id
	order by q.updated_at desc`

	query += " OFFSET ? "
	if offset > 0 {
		query += " FETCH NEXT 50 ROWS only;"
	} else {
		query += " limit 50;"
	}
	num, err := o.Raw(query, offset).Values(&maps)

	// query += " OFFSET ?"
	// if offset > 0 {
	// 	query += " FETCH NEXT '" + strconv.Itoa(PAGINATION_MAXIMUM) + "' ROWS only;"
	// } else {
	// 	query += " limit '" + strconv.Itoa(PAGINATION_MAXIMUM) + "';"
	// }
	// num, err := o.Raw(query, offset).Values(&maps)
	if err != nil {
		beego.Error(err)
		eMsg = err.Error()
	}
	res = DBResponse{num, eMsg, maps}
	return res
}

func QuestionAdd(campaignId, categoryId, questionText, questionImage, attempt, points, level, vFrom, vTo, status, isCampaign string) DBResponse {
	var maps []orm.Params
	eMsg := ""
	o := orm.NewOrm()
	query := "INSERT INTO tbl_question " +
		"(campaign_id, category_id, question_text, question_image, attempt_in_sec, points, level, valid_from, valid_to, status, created_by, created_at, updated_by, updated_at, is_campaign) " +
		"VALUES(?,?,?,?,?,?,?, "

	sD, errSD := ConvertToTime(vFrom)
	if errSD == nil {
		query += "'" + sD.Format("2006-01-02 15:04:05-07:00") + "', "
	} else {
		query += " null,"
	}
	eD, errED := ConvertToTime(vTo)
	if errED == nil {
		query += "'" + eD.Format("2006-01-02 15:04:05-07:00") + "', "
	} else {
		query += " null,"
	}

	query += "?,'admin',NOW(),'admin',NOW()"

	if isCampaign != "" {
		query += ",'" + isCampaign + "') RETURNING question_id; "
	} else {
		query += ") RETURNING question_id;"
	}

	num, err := o.Raw(query, campaignId, categoryId, questionText, questionImage, attempt, points, level, status).Values(&maps)
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

func QuestionGetById(question_id string) DBResponse {
	var maps []orm.Params
	eMsg := ""
	var res DBResponse
	o := orm.NewOrm()
	query := `select * from tbl_question where question_id = ? limit 1;`

	num, err := o.Raw(query, question_id).Values(&maps)
	if err != nil {
		beego.Error(err)
		eMsg = err.Error()
	}
	res = DBResponse{num, eMsg, maps}
	return res
}

func QuestionEdit(question_id, campaignId, category_id, question_text, question_image, attempt, points, level, vFrom, vTo, status, isCampaign string) DBResponse {
	eMsg := ""
	var num int64 = 0
	o := orm.NewOrm()
	query := "UPDATE tbl_question SET "
	sD, errSD := ConvertToTime(vFrom)
	if errSD == nil {
		query += " valid_from = '" + sD.Format("2006-01-02 15:04:05-07:00") + "', "
	} else {
		query += " valid_from =null,"
	}
	eD, errED := ConvertToTime(vTo)
	if errED == nil {
		query += " valid_to = '" + eD.Format("2006-01-02 15:04:05-07:00") + "', "
	} else {
		query += " valid_to =null,"
	}
	if question_image != "" {
		query += "question_image = '" + question_image + "',"
	}
	if isCampaign != "" {
		query += "is_campaign = '" + isCampaign + "',"
	}
	query += " campaign_id = ?, category_id = ?, question_text = ?, attempt_in_sec = ?, points = ?, level = ?,status = ?, created_by = 'admin', updated_by = 'admin', updated_at = now() "

	query += " where question_id = ?;"

	res, err := o.Raw(query, campaignId, category_id, question_text, attempt, points, level, status, question_id).Exec()
	if err != nil {
		beego.Error(err)
		eMsg = err.Error()
	} else {
		num, _ := res.RowsAffected()
		beego.Info("Total Rows affected : ", num)
	}
	return DBResponse{num, eMsg, nil}
}
