package router

import (
	"message/controller"

	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	//顯示 debug 模式
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	api := r.Group("api")
	{
		//新增留言
		api.POST("/message", controller.Create)
		//查詢全部留言
		api.GET("/message", controller.GetAll)
		//查詢 {id} 留言
		api.GET("/message/:id", controller.Get)
		//修改 {id} 留言
		api.PATCH("/message/:id", controller.Update)
		//刪除 {id} 留言
		api.DELETE("/message/:id", controller.Delete)

		//新增上傳檔案
		api.POST("/file", controller.CreateFile)
		//查詢全部上傳檔案
		api.GET("/file", controller.GetAllFile)
		//查詢 {id} 上傳檔案
		api.GET("/file/:id", controller.GetFile)
		//修改 {id} 上傳檔案
		api.PATCH("/file/:id", controller.UpdateFile)
		//刪除 {id} 上傳檔案
		api.DELETE("/file/:id", controller.DeleteFile)
	}

	return r
}
