package routers

import (
	"github.com/gin-gonic/gin"
	"gin-learning/controllers"
)

func Register() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("templates/**/*")

	articles := new(controllers.Articles)

	v1 := r.Group("/v1")
	{
		v1.GET("/index", articles.Index)
		v1.GET("/create", articles.Create)
		v1.GET("/update", articles.Update)
		v1.GET("/delete", articles.Delete)
		v1.POST("/store", articles.Store)
	}

	return r
}
