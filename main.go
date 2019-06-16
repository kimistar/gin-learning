package main

import (
	"context"
	"flag"
	"gin-learning/models"
	"gin-learning/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var env string
	flag.StringVar(&env, "env", "develop", "Input app environment develop/testing/product")
	flag.Parse()
	if env != "develop" && env != "testing" && env != "product" {
		os.Exit(1)
	}

	gin.SetMode(gin.ReleaseMode)
	// 注册路由
	r := routers.Register()
	// 建立连接池
	CreatePool(env)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 优雅重启
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Println("timeout of 5 seconds.")

	log.Println("Server exiting")
}

func CreatePool(env string) {
	models.Env = env
	models.ConnectDB()
	models.ConnectRedis()
}
