package main

import (
	"fmt"
	//"message/model"
	"message/model"
	"message/router"
	"message/sql"
)

func main() {
	//連線資料庫
	if err := sql.InitMySql(); err != nil {
		panic(err)
	}

	//連結模型
	sql.Connect.AutoMigrate(&model.Message{}, &model.UploadImage{})
	sql.Connect.Table("message")
	sql.Connect.Table("uploadimage")

	//註冊路由
	r := router.SetRouter()

	//啟動埠為8082的專案
	fmt.Println("開啟:8082...")
	r.Run(":8082")
}
