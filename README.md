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

想要获取请求体中的JSON数据，就要使用结构体，保证结构体的属性包含请求体内的所有参数。

```go
package main

import (
	"github.com/gin-gonic/gin"
)

// UserRequest 接收POST请求的结构体
type UserRequest struct {
	// json表示请求中的变量名
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()
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

	r.Run(":8080")
}
```

这样，就能获取到JSON数据。

![image-20260315115254494](README_Picture/image-20260315115254494.png)

### 1.2.3 XML传值

xml传值前，需要为结构体添加xml的Tag，才能识别xml的变量。

```go
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
```

![image-20260315120746005](README_Picture/image-20260315120746005.png)

## 1.3 路由分组

如果将所有的路由使用`r.GET`、`r.POST`等配置在main函数中，会导致代码复杂。因此需要对路由进行分组。

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//路由分组
	defaultRouters := r.Group("/")
	{
		// 获取分组后的路由来拼接路径
		defaultRouters.GET("/", func(c *gin.Context) {
			c.String(200, "首页")
		})
		defaultRouters.GET("/news", func(c *gin.Context) {
			c.String(200, "新闻")
		})
	}
	// 路由分组
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "api首页",
			})
		})
		apiRouters.GET("/news", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "api新闻",
			})
		})
	}
	r.Run(":8080")
}
```

第一个路由组的路径是`localhost:8080/`，第二个路由组是`localhost:8080/api/`。而路由组下的路由是根据这些路径来继续拼接的。

这样，可以将不同的路由放到不同的文件中分开维护。

`main.go`。

```go
package main

import (
	"Golang-learn/Demo3/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.InitDefaultRouters(r)
	routers.InitApiRouters(r)
	r.Run(":8080")
}
```

`routers/apiRouters.go`。

```go
package routers

import "github.com/gin-gonic/gin"

// InitApiRouters 初始化API路由
func InitApiRouters(r *gin.Engine) {
	// 路由分组
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "api首页",
			})
		})
		apiRouters.GET("/news", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "api新闻",
			})
		})
	}
}
```

`routers/defaultRouters.go`。

```go
package routers

import "github.com/gin-gonic/gin"

// InitDefaultRouters 初始化默认路由
func InitDefaultRouters(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		// 获取分组后的路由来拼接路径
		defaultRouters.GET("/", func(c *gin.Context) {
			c.String(200, "首页")
		})
		defaultRouters.GET("/news", func(c *gin.Context) {
			c.String(200, "新闻")
		})
	}
}
```

这样就实现了路由分组，让main文件显得干净易维护。

## 1.4 自定义控制器

Gin本身并没有控制器，但能够使用普通函数来完成类似Java内的Controller的功能。

在一般的项目中，通常会将监听路径`c.GET`与触发的函数分离。

```go
c.GET("/login", Login) // Login方法在其他文件实现
```

将上面的路由分组的触发函数进行分开。

`controllers/default/defaultController.go`。

```go
package _default

import "github.com/gin-gonic/gin"

// ShowHomePage 加载首页，对应defaultRouter.GET的"/"路径
func ShowHomePage(c *gin.Context) {
	c.String(200, "首页")
}

// ShowNewsPage 加载新闻，对应defaultRouter.GET的"/news"路径
func ShowNewsPage(c *gin.Context) {
	c.String(200, "新闻")
}
```

`controllers/api/apiControllers.go`。

```go
package api

import "github.com/gin-gonic/gin"

// ApiShowHomePage 加载首页，对应apiRouter.GET的"/"路径
func ApiShowHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "api 首页",
	})
}

// ApiShowNewsPage 对应apiRouter.GET的"/news"路径
func ApiShowNewsPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "api 新闻",
	})
}
```

`routers/apiRouters.go`。

```go
package routers

import (
	"Golang-learn/Demo4_controllers/controllers/api"

	"github.com/gin-gonic/gin"
)

// InitApiRouters 初始化API路由
func InitApiRouters(r *gin.Engine) {
	// 路由分组
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", api.ApiShowHomePage)
		apiRouters.GET("/news", api.ApiShowNewsPage)
	}
}
```

`routers/defaultRouters.go`。

```go
package routers

import (
	"Golang-learn/Demo4_controllers/controllers/_default"

	"github.com/gin-gonic/gin"
)

// InitDefaultRouters 初始化默认路由
func InitDefaultRouters(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		// 获取分组后的路由来拼接路径
		defaultRouters.GET("/", _default.ShowHomePage)
		defaultRouters.GET("/news", _default.ShowNewsPage)
	}
}
```

`main.go`。

```go
package main

import (
	"Golang-learn/Demo3/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.InitDefaultRouters(r)
	routers.InitApiRouters(r)
	r.Run(":8080")
}
```

这样，就成功实现routers包只实现路由转发，在controllers下执行触发路由逻辑。

但这种情况下，这些方法依旧是散的，只能执行，没办法让其他的类也能够使用。因此可以构建一个结构体，让这些方法成为结构体的方法，那么直接使用结构体就能实现方法的使用和继承，且更结构化。

以api下的`apiController.go`为例。

```go
package api

import (
	"Golang-learn/Demo4_controllers/controllers/_default"

	"github.com/gin-gonic/gin"
)

type ApiController struct {
	_default.DefaultController
}

// ShowHomePage 加载首页，对应apiRouter.GET的"/"路径
func (ApiController *ApiController) ShowHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "api 首页",
	})
}

// ShowNewsPage 对应apiRouter.GET的"/news"路径
func (ApiController *ApiController) ShowNewsPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "api 新闻",
	})
}
```

那么在`apiRouters.go`中的调用可以改为如下。

```go
package routers

import (
	"Golang-learn/Demo4_controllers/controllers/api"

	"github.com/gin-gonic/gin"
)

// InitApiRouters 初始化API路由
func InitApiRouters(r *gin.Engine) {
	var apiController api.ApiController

	// 路由分组
	apiRouters := r.Group("/api")
	{
		// 临时创建结构体实例进行使用
		apiRouters.GET("/", api.ApiController{}.ShowHomePage)
		// 先构建实例，再调用方法
		apiRouters.GET("/news", apiController.ShowNewsPage)
		// 继承DefaultController的方法
		apiRouters.GET("/default", apiController.ShowDefaultHomePage)
	}
}
```

这样即能够让调用方法的方式结构化，而且由于结构体已经表示了当前方法的归属，因此方法名称不需要写为`ApiShowHomePage`，直接写为`ShowHomePage`即可。

而且ApiController结构体继承了DefaultController结构体，那么使用ApiController实例能够直接调用父类的方法。

## 1.5 中间件

中间件就是在匹配路由前和匹配路由后执行的函数。

```go
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
	fmt.Println("Hello World")
}
```

上面实现了基础的中间件。这样，在获取请求前，先回执行局部中间件，然后执行全局中间件，最后才执行请求。

![image-20260315183137300](README_Picture/image-20260315183137300.png)

如果使用`c.Next`，就说明跳过了下面的代码，先执行请求，请求执行完毕后回到这个函数，继续执行下面的代码。

如果使用`c.Abort()`，那么当前中间件执行完毕后，就不会继续向后执行，而是直接停止程序。

同样的，如果在路由组中使用中间件，就只有当前的路由组会执行中间件，其他组不会。
