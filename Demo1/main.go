package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建路由引擎
	r := gin.Default()
	// String响应
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	// JSON响应
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	// JSONP响应
	r.GET("/jsonp", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	// xml响应
	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	// html响应
	// 指定模板的目录
	r.LoadHTMLGlob("templates/*")
	r.GET("/html", func(c *gin.Context) {
		// 加载模板，此时gin.H用于指定模板里的变量
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "Hello World",
		})
	})
	r.Run(":8080")
}
