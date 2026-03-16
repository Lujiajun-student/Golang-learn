package main

import (
	"Golang-learn/Demo5_middleware/middle"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 全局使用中间件
	r.Use(middle.InitMiddleware)
	// 针对某个请求使用中间件
	r.GET("/", func(c *gin.Context) {
		c.String(200, "首页")
	}, middle.PrintHello)

	r.Run(":8080")
}
