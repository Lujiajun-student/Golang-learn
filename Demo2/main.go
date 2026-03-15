package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 在路径中设置age参数，能够捕获参数
	r.GET("/:age", func(c *gin.Context) {
		// 在上下文中添加参数
		c.Set("username", "小张")
		// 从上下文中获取参数
		username, _ := c.Get("username")
		// 从URL中获取参数，如果没有则使用默认值
		age := c.DefaultQuery("age", "30")
		c.JSON(200, gin.H{
			"username": username,
			"age":      age,
		})
	})

	// POST传值
	r.POST("/login", func(c *gin.Context) {
		id := c.DefaultPostForm("id", "0")
		password := c.PostForm("password")
		c.JSON(200, gin.H{
			"id":       id,
			"password": password,
		})
	})

	r.Run(":8080")
}
