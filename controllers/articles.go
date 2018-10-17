package controllers

import (
	"gin-learning/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Articles struct {
}

func (_ *Articles) Index(ctx *gin.Context) {
	articleModel := new(models.Articles)
	list := articleModel.List()
	ctx.HTML(http.StatusOK, "articles/index.html", gin.H{
		"list": list,
	})
}

func (_ *Articles) Create(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "articles/create-edit.html", nil)
}

func (_ *Articles) Edit(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Redirect(http.StatusFound, "/articles")
		return
	}
	articleModel := new(models.Articles)
	article := articleModel.First(id)
	ctx.HTML(http.StatusOK, "articles/create-edit.html", gin.H{
		"article": article,
	})
}

func (_ *Articles) Store(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	title := ctx.PostForm("title")
	author := ctx.PostForm("author")
	content := ctx.PostForm("content")
	articleModel := new(models.Articles)
	if id == 0 {
		articleModel.Insert(title, author, content)
	} else {
		articleModel.Edit(id, title, author, content)
	}

	ctx.Redirect(http.StatusFound, "/articles")
}

func (_ *Articles) Del(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Redirect(http.StatusFound, "/articles")
		return
	}
	articleModel := new(models.Articles)
	articleModel.Del(id)
	ctx.Redirect(http.StatusFound, "/articles")
}
