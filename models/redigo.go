package models

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
	"strconv"
	"time"
)

var redisPool *redis.Pool

func ConnectRedis() {
	// 主机
	host := GetConfig("redis.host")
	// 端口
	port := GetConfig("redis.port")
	// 密码
	password := GetConfig("redis.password")
	// 最大空闲连接数
	maxIdleConns, _ := strconv.Atoi(GetConfig("redis.max_idle_conns"))
	// 最大打开的连接数
	maxOpenConns, _ := strconv.Atoi(GetConfig("redis.max_open_conns"))

	redisPool = &redis.Pool{
		MaxIdle:     maxIdleConns,
		MaxActive:   maxOpenConns,
		IdleTimeout: 5 * time.Minute,
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
