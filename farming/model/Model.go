package model

import "gorm.io/gorm"

func (Message) TableName() string {
	return "message"
}

type Message struct {
	Id      int    `gorm:"primary_key,type:INT;not null;AUTO_INCREMENT"`
	User_Id int    `json:"User_Id"  binding:"required"`
	Content string `json:"Content"  binding:"required"`
	Version int    `gorm:"default:0"`
	// 包含 CreatedAt 和 UpdatedAt 和 DeletedAt 欄位
	gorm.Model
}

func (UploadImage) TableName() string {
	return "uploadimage"
}

// 設定 uploadImage 紀錄
type UploadImage struct {
	Id         int    `gorm:"primary_key,type:INT;not null;AUTO_INCREMENT"`
	Sender     string `json:"Sender"  binding:"required"`
	Location   string `json:"Location" `
	Machine_Id string `json:"Machine_Id" `
	Image      string `json:"Image" `

	// 包含 CreatedAt 和 UpdatedAt 和 DeletedAt 欄位
	gorm.Model
}
