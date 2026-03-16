package main

import (
	"Golang-learn/Demo9_gorm/models"
	"Golang-learn/Demo9_gorm/router/adminRouter"
	"Golang-learn/Demo9_gorm/router/defaultRouter"
	"Golang-learn/Demo9_gorm/router/userRouter"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 初始化数据库
	models.InitMySQL()
	// 配置session中间件
	// cookie.NewStore 表示将session保存到浏览器的Cookie中,密钥选择为secret123
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", "", []byte("secret123"))
	// 设置Session在Cookie的保存名称，可通过sessions.Default(c)来获取这个session对象
	router.Use(sessions.Sessions("mysession", store))

	userRouter.InitUserRouter(router)
	adminRouter.InitAdminRouter(router)
	defaultRouter.InitDefaultRouter(router)
	router.Run(":8080")
}
