package main

import (
	"fmt"
	"os"
	"gin-learning/routers"
	"github.com/go-ini/ini"
	"github.com/gin-gonic/gin"
	_ "gin-learning/orms"
)

func main() {
	// 加载配置
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// 运行模式
	mode := cfg.Section("").Key("app_mode").String()

	if mode == "product" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// 注册路由
	r := routers.Register()

	// 加载静态文件
	r.Static("/static", "static")

	http_port := cfg.Section("").Key("http_port").String()

	r.Run(http_port);
}
