package controller

import (
	"message/model"
	"message/repository"
	"net/http"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	message, err := repository.GetAllMessage()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func Get(c *gin.Context) {
	var message model.Message

	if err := repository.GetMessage(&message, c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "找不到留言"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func Create(c *gin.Context) {
	var message model.Message

	if c.PostForm("Content") == "" || utf8.RuneCountInString(c.PostForm("Content")) >= 20 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "沒有輸入內容或長度超過20個字元"})
		return
	}

	c.Bind(&message)
	repository.CreateMessage(&message)
	c.JSON(http.StatusCreated, gin.H{"message": message})
}

func Update(c *gin.Context) {
	var message model.Message

	if c.PostForm("Content") == "" || utf8.RuneCountInString(c.PostForm("Content")) >= 20 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "沒有輸入內容或長度超過20個字元"})
		return
	}

	if err := repository.UpdateMessage(&message, c.PostForm("Content"), c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "找不到留言"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func Delete(c *gin.Context) {
	var message model.Message

	if err := repository.DeleteMessage(&message, c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "找不到留言"})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "刪除留言成功"})
}

func CreateFile(c *gin.Context) {
	var uploadimage model.UploadImage

	if c.PostForm("Sender") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"uploadimage": "請輸入 傳送人"})
		return
	}

	// 接收並儲存檔案
	file, err := c.FormFile("Image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "檔案上傳錯誤: " + err.Error()})
		return
	}

	path := "uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "檔案儲存錯誤: " + err.Error()})
		return
	}

	// 將檔案路徑設定到 uploadimage 結構體的 Image 屬性
	uploadimage.Image = path

	c.Bind(&uploadimage)
	repository.CreateFile(&uploadimage)
	c.JSON(http.StatusCreated, gin.H{"uploadimage": uploadimage})
}

func GetAllFile(c *gin.Context) {
	uploadimage, err := repository.GetAllFile()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"uploadimage": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"uploadimage": uploadimage})
}

func GetFile(c *gin.Context) {
	var uploadimage model.UploadImage

	if err := repository.GetFile(&uploadimage, c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"uploadimage": "找不到上傳紀錄"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"uploadimage": uploadimage})
}

func UpdateFile(c *gin.Context) {
	var uploadimage model.UploadImage

	if c.PostForm("Sender") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"uploadimage": "請輸入 傳送人"})
		return
	}

	// 接收並儲存檔案 (不開放修改影片/圖片)
	// file, err := c.FormFile("Image")
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "檔案上傳錯誤: " + err.Error()})
	// 	return
	// }

	// path := "uploads/" + file.Filename
	// if err := c.SaveUploadedFile(file, path); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "檔案儲存錯誤: " + err.Error()})
	// 	return
	// }

	// 將檔案路徑設定到 uploadimage 結構體的 Image 屬性
	// uploadimage.Image = path

	if err := repository.UpdateFile(&uploadimage, c.PostForm("Sender"), c.PostForm("Machine_Id"), c.PostForm("Location"), c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "找不到上傳紀錄"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"uploadimage": uploadimage})
}

func DeleteFile(c *gin.Context) {
	var uploadimage model.UploadImage

	if err := repository.DeleteFile(&uploadimage, c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "找不到上傳紀錄"})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "刪除上傳紀錄成功"})
}
