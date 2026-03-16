package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 指定最大上传文件大小为8MB
	router.MaxMultipartMemory = 8 << 20
	// 指定静态页面的位置
	router.LoadHTMLGlob("templates/*")
	// 指定上传接口，用来展示上传图片的页面
	router.GET("/upload", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})
	// 点击页面的提交按钮后，执行下面的请求
	router.POST("/doUpload", func(c *gin.Context) {
		// 从上下文中读取上传的文件
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			log.Println(file.Filename)
			// 将文件保存到当前项目的upload目录下
			err := c.SaveUploadedFile(file, "./upload/"+file.Filename)
			if err != nil {
				c.String(400, "failed to save file")
				return
			}
		}
		c.String(200, "上传成功")
	})
	router.Run(":8080")
}
