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

func (article *Articles) Insert(title,author,content string) bool {
	createTime,_:=time.Parse("2006-01-02 15:04:05",time.Now().Format("2006-01-02 15:04:05"))
	article=&Articles{Title:title,Author:author,Content:content,CreateTime:createTime}
	orm.Create(article)
	if orm.NewRecord(article) {
		return false
	}
	return true
}

func (article *Articles) Edit (id int,title,author,content string) bool {
	ret:=article.First(id)
	if ret == nil {
		return false
	}
	updateTime,_:=time.Parse("2006-01-02 15:04:05",time.Now().Format("2006-01-02 15:04:05"))
	orm.Model(ret).Updates(map[string]interface{}{"title":title,"author":author,"content":content,"update_time":updateTime})
	return true
}

func (article *Articles) Del(id int) bool {
	ret:=article.First(id)
	if ret == nil {
		return false
	}
	orm.Delete(ret)
	return true
}
