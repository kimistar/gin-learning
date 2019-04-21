package routers

import (
	"core"
	"gin-learning/controllers"
	"gin-learning/middleware"
	"github.com/gin-gonic/gin"
)

func Register() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/test", core.Handle(controllers.Test))

	articles := new(controllers.Articles)

	v1 := r.Group("/")
	v1.Use(core.Handle(middleware.Auth))
	{
		v1.GET("/articles", core.Handle(articles.Index))
		v1.GET("/article/create", core.Handle(articles.Create))
		v1.GET("/article/edit/:id", core.Handle(articles.Edit))
		v1.GET("/article/del/:id", core.Handle(articles.Del))
		v1.POST("/article/store", core.Handle(articles.Store))
	}

	return r
}
