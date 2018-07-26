package controller

/*
首页的控制器
*/

import (
	"net/http"
	"log"
	"app/model"
)

func Index(w http.ResponseWriter, r *http.Request)string {


	//var boss public.Mysql.DBModel = public.Mysql.DBModel{ Table:"boss" }

	//var boss public.Mysql.

	//log.Print(boss)
	//var user model.DBModel = model.DBModel{ Table:"users" }

	keyArr := make(map[string]interface{})

	keyArr["id"] = "boss_333"
	keyArr["name"] = "333A"
	keyArr["ph"] = 33
	keyArr["pw"] = 33
	keyArr["ch"] = 33
	keyArr["resoures"] = "321321"
	keyArr["per"] = 41
	//
	model.Boss.New(keyArr, true).Save()	// 增加一行新数据
	res, err := model.User.Where("user_id='f639a58a36c1c'" , "OR", "user_id='b8255c8370dc0'").Del()

	log.Print( err  ) // 删除指定一行
	log.Print( res.LastInsertId()) // 删除指定一行
	log.Print( res.RowsAffected()  ) // 删除指定一行

	str := "hello world!"
	return str
}
