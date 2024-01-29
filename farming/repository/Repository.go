package repository

import (
	"message/model"
	"message/sql"
)

//查詢全部留言
func GetAllMessage() (message []*model.Message, err error) {
	err = sql.Connect.Find(&message).Error
	return
}

//查詢 {id} 留言
func GetMessage(message *model.Message, id string) (err error) {
	err = sql.Connect.Where("id=?", id).First(&message).Error
	return
}

//新增留言
func CreateMessage(message *model.Message) (err error) {
	err = sql.Connect.Create(&message).Error
	return
}

//更新 {id} 留言
func UpdateMessage(message *model.Message, content, id string) (err error) {
	err = sql.Connect.Where("id=?", id).First(&message).Update("content", content).Error
	return
}

//刪除 {id} 留言
func DeleteMessage(message *model.Message, id string) (err error) {
	err = sql.Connect.Where("id=?", id).First(&message).Delete(&message).Error
	return
}

// 新增照片/檔案
func CreateFile(uploadimage *model.UploadImage) (err error) {
	err = sql.Connect.Create(&uploadimage).Error
	return
}

// 查詢全部上傳照片記錄
func GetAllFile() (uploadimage []*model.UploadImage, err error) {
	err = sql.Connect.Find(&uploadimage).Error
	return
}

//查詢 {id} 上傳照片
func GetFile(uploadimage *model.UploadImage, id string) (err error) {
	err = sql.Connect.Where("id=?", id).First(&uploadimage).Error
	return
}

// 更新 {id} 照片
func UpdateFile(uploadimage *model.UploadImage, Sender, Machine_Id, Location, id string) (err error) {
	err = sql.Connect.Where("id=?", id).First(&uploadimage).Updates(model.UploadImage{
		Sender:     Sender,
		Machine_Id: Machine_Id,
		Location:   Location,
	}).Error

	return
}

//刪除 {id} 照片/影片
func DeleteFile(uploadimage *model.UploadImage, id string) (err error) {
	err = sql.Connect.Where("id=?", id).First(&uploadimage).Delete(&uploadimage).Error
	return
}
