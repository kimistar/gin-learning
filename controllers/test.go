package controllers

import (
	"core"
	"github.com/gin-gonic/gin"
)

func Test(c *core.Context) {
	c.Success(gin.H{
		"name":       "lilei",
		"girlfriend": "hanmeimei",
	})
}
