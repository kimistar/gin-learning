package models

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/gomodule/redigo/redis"
	"os"
	"time"
)

var redisPool *redis.Pool

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
	host := cfg.Section(mode).Key("redis.host").String()
	// 端口
	port := cfg.Section(mode).Key("redis.port").String()
	// 密码
	password := cfg.Section(mode).Key("redis.password").String()
	// 最大空闲连接数
	maxIdleConns, err = cfg.Section(mode).Key("redis.max_idle_conns").Int()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	// 最大打开的连接数
	maxOpenConns, err = cfg.Section(mode).Key("redis.max_open_conns").Int()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	redisPool = &redis.Pool{
		MaxIdle:     maxIdleConns,
		MaxActive:   maxOpenConns,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host+":"+port)
			if err != nil {
				fmt.Printf("%v", err)
				os.Exit(1)
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					fmt.Printf("%v", err)
					os.Exit(1)
				}
			}
			return c, nil
		},
	}
}

func GetRedisPool() *redis.Pool {
	return redisPool
}
