package controllers

import (
	"gin-learning/core"
	"gin-learning/models"
	"github.com/gin-gonic/gin"
)

type TestReq struct {
	ID     string `json:"id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Age    int    `json:"age" binding:"required"`
	School string `json:"school" binding:"required"`
}

func Demo(c *core.Context) {
	var req TestReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Fail(10001, "缺少参数", nil)
		return
	}

	pupil := &models.Pupil{}
	if !pupil.HasACar(req.Name) {
		c.Fail(10002, "NO", nil)
		return
	}

	c.Success(gin.H{
		"id": req.ID,
	})
}
