package models

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"time"
)

var orm *gorm.DB

func init() {
	var err error
	var cfg *ini.File
	var maxIdleConns int
	var maxOpenConns int

	// load配置
	cfg, err = ini.Load("conf/database.ini", "conf/app.ini")
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	// 运行模式
	mode := cfg.Section("").Key("app_mode").String()
	// 主机
	host := cfg.Section(mode).Key("mysql.host").String()
	// 端口
	port := cfg.Section(mode).Key("mysql.port").String()
	// 用户名
	username := cfg.Section(mode).Key("mysql.username").String()
	// 密码
	password := cfg.Section(mode).Key("mysql.password").String()
	// 数据库名称
	dbname := cfg.Section(mode).Key("mysql.dbname").String()
	// 最大空闲连接数
	maxIdleConns, err = cfg.Section(mode).Key("mysql.max_idle_conns").Int()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	// 最大打开的连接数
	maxOpenConns, err = cfg.Section(mode).Key("mysql.max_open_conns").Int()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=true&loc=Local"

	orm, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Fail to open mysql: %v", err)
		os.Exit(1)
	}

	orm.DB().SetMaxIdleConns(maxIdleConns)
	orm.DB().SetMaxOpenConns(maxOpenConns)
	orm.DB().SetConnMaxLifetime(time.Hour)
}

func GetGorm() *gorm.DB {
	return orm
}
