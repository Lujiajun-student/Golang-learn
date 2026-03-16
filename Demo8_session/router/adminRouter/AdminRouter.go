package adminRouter

import (
	"Golang-learn/Demo8_session/controller/admin"

	"github.com/gin-gonic/gin"
)

func InitAdminRouter(r *gin.Engine) {
	adminGroup := r.Group("/admin")
	{
		adminGroup.GET("/index", admin.AdminIndex)
		adminGroup.GET("/news", admin.AdminNews)
	}
}
