package main

import (
    // "net/http"
	"fmt" 
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.POST("/upload", func(c *gin.Context) {
        file, _ := c.FormFile("file")
        // 示例：打印文件名
        if file != nil {
            fmt.Println("Uploaded file:", file.Filename)
        }

        // 提取其他表单数据
        time := c.PostForm("time")
        location := c.PostForm("location")
        sender := c.PostForm("sender")
        machineId := c.PostForm("machineId")

        // 示例：打印表单数据
        fmt.Println("Time:", time)
        fmt.Println("Location:", location)
        fmt.Println("Sender:", sender)
        fmt.Println("Machine ID:", machineId)

        // 可以在这里添加对这些数据的处理
        // 例如存储到 Azure SQL Database

        c.JSON(200, gin.H{"status": "file uploaded"})
    })

    router.Run(":8083")
}

// type TestData struct {
//     Hello string `json:"hello"`
// }

// func test(c *gin.Context) {
//     data := new(TestData)
//     data.Hello = "world!"
//     c.JSON(http.StatusOK, data)
// }

// func main() {
//     server := gin.Default()
//     server.GET("/test", test)
//     server.Run(":8083")
// }