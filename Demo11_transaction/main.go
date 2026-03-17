package main

import (
	"Golang-learn/Demo11_transaction/models"
	"Golang-learn/Demo11_transaction/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.InitMySQL()
	router.InitUserRouter(r)
	r.Run()
}
