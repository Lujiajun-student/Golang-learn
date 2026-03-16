package userRouter

import (
	"Golang-learn/Demo8_session/controller/user"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.GET("/info", user.UserInfo)

		userGroup.GET("/news", user.UserNews)
	}
}
