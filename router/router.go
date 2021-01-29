package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-greeting/pkg/e"
	"github.com/go-greeting/pkg/setting"
	v1 "github.com/go-greeting/router/api/v1"
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

	apiV1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiV1.GET("/tags", v1.GetTags)
		//新建标签
		apiV1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiV1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiV1.DELETE("/tags/:id", v1.DeleteTag)

		apiV1.GET("/articles", v1.GetArticles)
		apiV1.GET("/articles/:id", v1.GetArticle)
		apiV1.POST("/articles", v1.AddArticle)
		apiV1.PUT("articles/:id", v1.EditArticle)
		apiV1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
