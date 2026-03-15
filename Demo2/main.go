package main

import (
	"encoding/xml"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserRequest 接收POST请求的结构体
type UserRequest struct {
	// json表示请求中的变量名
	Username string `json:"username" xml:"username"`
	Password string `json:"password" xml:"password"`
}

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

	// POST的JSON传值
	r.POST("/user", func(c *gin.Context) {
		var req UserRequest
		// 使用ShouldBind绑定JSON数据到结构体
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		// 获取到包含参数的结构体，可以返回到页面上
		c.JSON(200, gin.H{
			"username": req.Username,
			"password": req.Password,
		})

	})
	// POST的XML传值
	r.POST("/user/xml", func(c *gin.Context) {
		// 获取xml数据
		b, _ := c.GetRawData()
		user := &UserRequest{}
		// 通过反序列化，将xml数据绑定到结构体上
		if err := xml.Unmarshal(b, user); err == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
		}
	})

	r.Run(":8080")
}
