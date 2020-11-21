package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Moqqll/go_webserver/dbops/mysql"
	"github.com/Moqqll/go_webserver/dbops/redis"
	"github.com/Moqqll/go_webserver/logger"
	"github.com/Moqqll/go_webserver/routes"
	"github.com/Moqqll/go_webserver/setting"

	"go.uber.org/zap"
)

func main() {
	//加载配置
	if err := setting.Init(); err != nil {
		fmt.Printf("init settings failde, err:%v\n", err)
	}

	//初始化日志
	if err := logger.Init(setting.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failde, err:%v\n", err)
	}
	defer zap.L().Sync()

	//初始化mysql连接
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failde, err:%v\n", err)
	}
	defer mysql.Close()

	//初始化redis连接
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redix failde, err:%v\n", err)
	}
	defer redis.Close()

	//初始化ID生成器

	//注册路由
	r := routes.Setup()

	//设置启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", setting.Conf.Port),
		Handler: r,
	}

	go func() {
		//开启一个goroutine来启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Error("listen:%s\n", zap.Error(err))
		}
	}()

	//等待中断信号来优雅地关闭服务器，为关闭服务器设置一个5秒的超时
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) //此处不会阻塞
	<-quit                                               //阻塞在此，当接收到上述两种信号时才会往下执行

	zap.L().Info("Shutdown server ...")
	//创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//优雅关机：5秒内优雅关闭服务器（待未处理完的请求处理完再关闭服务器），超过5秒则直接退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("Server Shutdown:%s\n", zap.Error(err))
	}
	zap.L().Info("Server exiting.")
}
