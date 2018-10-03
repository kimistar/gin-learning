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
	orm.Where(&Articles{ID:id}).First(article)
	return article
}

// 获取文章列表
func (article *Articles) List() []*Articles  {
	var articles []*Articles
	orm.Select("id,title,author,content,click,create_time").Order("id desc").Find(&articles)
	return articles
}
