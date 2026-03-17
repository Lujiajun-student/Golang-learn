package router

import (
	"Golang-learn/Demo11_transaction/controller"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.GET("/info", controller.UserController{}.UserInfo)
	}
}
