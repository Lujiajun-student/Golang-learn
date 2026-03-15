package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 全局使用中间件
	r.Use(PrintHello)
	// 针对某个请求使用中间件
	r.GET("/", func(c *gin.Context) {
		c.String(200, "首页")
	}, PrintHello)

	r.Run(":8080")
}

// PrintHello 定义需要的中间件函数
func PrintHello(c *gin.Context) {
	c.Next()
	fmt.Println("Hello World")
}
