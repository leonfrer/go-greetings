package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-greeting/pkg/e"
	"github.com/go-greeting/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(e.SUCCESS, gin.H{
			"massage": "test",
		})
	})

	return r
}
