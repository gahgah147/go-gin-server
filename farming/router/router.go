package router

import (
	"message/controller"
	"message/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Swagger 註解

// @title 智慧農業專案 API
// @version 1.0
// @description 智慧農業專案 API swagger 文件

// @contact.name nalson
// @contact.url https://gahgah147.github.io/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8082
// schemes http
func SetRouter() *gin.Engine {
	//顯示 debug 模式
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

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

	// 設定swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
