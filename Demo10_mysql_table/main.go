package main

import (
	"Golang-learn/Demo10_mysql_table/models"
	"Golang-learn/Demo10_mysql_table/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 初始化数据库
	models.InitMySQL()
	// 初始化路由
	router.InitArticleRouter(r)

	r.Run(":8080")
}
