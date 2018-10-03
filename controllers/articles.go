package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin-learning/models"
	"strconv"
)

type Articles struct {
}

func (this *Articles) Index(ctx *gin.Context) {
	articleModel := new(models.Articles)
	list:=articleModel.List()
	ctx.HTML(http.StatusOK, "articles/index.html", gin.H{
		"list":list,
	})
}

func (this *Articles) Create(ctx *gin.Context) () {
	ctx.HTML(http.StatusOK,"articles/create-edit.html",nil)
}

func (this *Articles) Edit(ctx *gin.Context) {
	id,err:=strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}
	articleModel:=new(models.Articles)
	article:=articleModel.First(id)
	ctx.HTML(http.StatusOK,"articles/create-edit.html",gin.H{
		"article":article,
	})
}

func (this *Articles) Store(ctx *gin.Context) {

}

func (this *Articles) Del(ctx *gin.Context) {

}
