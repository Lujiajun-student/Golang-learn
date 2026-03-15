# Golang-learn
Exercise from https://www.bilibili.com/video/BV1XY4y1t76G/?p=31&amp;spm_id_from=333.1007.top_right_bar_window_history.content.click

# 1. Gin框架

Gin最擅长Api接口的高并发。如果业务简单、规模不大，就使用Gin。

如果从一个新项目开始，就要执行go的初始化。

```cmd
go mod init Golang-learn
go get -u github.com/gin-gonic/gin
```

```go
package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建路由引擎
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})
	r.Run(":8080")
}
```

上面就是一个简单的Go项目。

![image-20260315105748849](README_Picture/image-20260315105748849.png)

## 1.1 响应数据

获取到`gin.Default()`后，有多种方法能够创建路由并返回数据。比如在页面中返回`c.JSON()`，浏览器会显示JSON数据；`c.String()`在浏览器显示字符串；`c.JSONP()`会显示包裹在回调函数内的JSON数据；`c.XML()`会显示xml数据，`c.HTML()`会在浏览器显示html数据。

![image-20260315110950539](README_Picture/image-20260315110950539.png)

## 1.2 GET和POST传值

### 1.2.1 Get传值

Get传值主要通过路径携带参数。而从路径上获取参数，只需要在GET请求的路径上添加`/:${value}`即可，其中`${value}`指变量名。这样，在请求内可通过`c.Query(value)`来获取该变量。

```go
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
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
		fmt.Println(username, age)
	})

	r.Run(":8080")
}
```

![image-20260315113208696](README_Picture/image-20260315113208696.png)

### 1.2.2 POST传值

POST请求的参数会保存在请求体中，通过`c.PostForm`和`c.DefaultPostForm`来获取。

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
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
```

值得注意的是，这种方法只能接收请求中的`form-data`，这样才能接收参数。如果请求体的参数以JSON的形式来传输，就无法接收。

![image-20260315113743444](README_Picture/image-20260315113743444.png)

