package middle

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(c *gin.Context) {
	fmt.Println(time.Now())
	fmt.Println("请求路径: ", c.Request.URL)
}

func PrintHello(c *gin.Context) {
	fmt.Println("Hello World")
}
