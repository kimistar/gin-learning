package routers

import (
	"github.com/gin-gonic/gin"
	"gin-learning/controllers"
)

func Register() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	articles := new(controllers.Articles)

	v1:=r.Group("/")
	{
		v1.GET("/articles", articles.Index)
		v1.GET("/article/create", articles.Create)
		v1.GET("/article/edit/:id", articles.Edit)
		v1.GET("/article/del/:id", articles.Del)
		v1.POST("/article/store", articles.Store)
	}

	return r
}
