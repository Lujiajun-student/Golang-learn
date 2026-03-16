package userRouter

import (
	"Golang-learn/Demo9_gorm/controller/user"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.GET("/info", user.UserInfo)

		userGroup.GET("/news", user.UserNews)

		userGroup.POST("/create", user.CreateUser)

		userGroup.PUT("/edit", user.EditUser)

		userGroup.DELETE("/delete", user.DeleteUser)

		userGroup.GET("/show", user.ShowUser)
	}
}
