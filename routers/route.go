package routers

import (
	"github.com/gin-gonic/gin"
	"test/pkg/setting"
)

func InitRoute() *gin.Engine {
	r := gin.Default()
	gin.SetMode(setting.RunMode)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "hello world"})
	})

	return r
}
