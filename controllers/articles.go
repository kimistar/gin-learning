package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin-learning/models"
)

type Articles struct {
}

func (this *Articles) Index(ctx *gin.Context) {
	//ctx.HTML(http.StatusOK, "articles/index.html", gin.H{
	//	"title": "测试文章",
	//})

	article := new(models.Articles)
	ret := article.First(1)

	ctx.JSON(http.StatusOK, gin.H{
		"id":     ret.ID,
		"title":  ret.Title,
		"author": ret.Author,
	})
}

func (this *Articles) Create(ctx *gin.Context) () {
	ctx.String(http.StatusOK, ctx.DefaultQuery("name", "create"))
}

func (this *Articles) Update(ctx *gin.Context) {
	ctx.String(http.StatusOK, "update");
}

func (this *Articles) Store(ctx *gin.Context) {

}

func (this *Articles) Delete(ctx *gin.Context) {

}
