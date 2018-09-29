package models

import (
	"time"
)

type Articles struct {
	ID         int
	Title      string
	Author     string
	Content    string
	Click      int
	CreateTime time.Time
	UpdateTime time.Time
}

// 用id查询一条记录
func (article *Articles) First(id int) *Articles {
	//orm.Where(&Articles{ID:id}).First(article)
	orm.Where("id = ?",id).First(article)
	return article
}
