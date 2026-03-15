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
