package routers

import (
	"gin-learning/controllers"
	"gin-learning/core"
	"gin-learning/middleware"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register() *gin.Engine {
	r := gin.New()
	// 加载模板文件
	r.LoadHTMLGlob("templates/**/*")
	// 加载静态文件
	r.Use(static.Serve("/", static.LocalFile("statics", false)))

	// 500
	r.Use(core.Handle(middleware.Recovery))
	// 404
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "error/404", nil)
	})

	v1 := r.Group("/")
	v1.Use(core.Handle(middleware.Auth))
	{
		articles := new(controllers.Articles)
		v1.GET("/articles", core.Handle(articles.Index))
		v1.GET("/article/create", core.Handle(articles.Create))
		v1.GET("/article/edit/:id", core.Handle(articles.Edit))
		v1.GET("/article/del/:id", core.Handle(articles.Del))
		v1.POST("/article/store", core.Handle(articles.Store))
	}

	r.POST("/demo", core.Handle(controllers.Demo))

	return r
}
