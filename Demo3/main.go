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
