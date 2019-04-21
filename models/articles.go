package models

import (
	"time"
)

type Articles struct {
	ID         int
	Title      string
	Author     string
	Content    string
	Thumbnail  string
	CreateTime time.Time
	UpdateTime time.Time
}

// 用id查询一条记录
func (article *Articles) First(id int) *Articles {
	orm.Where(&Articles{ID: id}).First(article)
	return article
}

// 获取文章列表
func (_ *Articles) List() []Articles {
	var articles []Articles
	orm.Select("id,title,author,content,thumbnail,create_time").Order("id desc").Find(&articles)
	return articles
}

// 返回数据插入成功后的ID
func (_ *Articles) Insert(title, author, content string) int {
	now := time.Now()
	article := &Articles{Title: title, Author: author, Content: content, CreateTime: now, UpdateTime: now}
	orm.Create(article)
	return article.ID
}

// 返回受影响行数
func (article *Articles) Edit(id int, title, author, content string) int64 {
	ret := article.First(id)
	// 查无结果 ret为空的Article
	if ret.ID == 0 {
		return 0
	}
	updateTime := time.Now()
	rowsAffected := orm.Model(ret).Updates(map[string]interface{}{"title": title, "author": author, "content": content, "update_time": updateTime}).RowsAffected
	return rowsAffected
}

// 返回受影响行数
func (article *Articles) Del(id int) int64 {
	ret := article.First(id)
	if ret.ID == 0 {
		return 0
	}
	rowsAffected := orm.Delete(ret).RowsAffected
	return rowsAffected
}
