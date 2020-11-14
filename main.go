package main

import (
	"fmt"
	"go_webserver/dbops/mysql"
	"go_webserver/dbops/redis"
	"go_webserver/logger"
	"go_webserver/route"
	"go_webserver/setting"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	//加载配置
	if err := setting.Init(); err != nil {
		fmt.Printf("init settings failde, err:%v\n", err)
	}

	//初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("init logger failde, err:%v\n", err)
	}
	defer zap.L().Sync()

	//初始化mysql连接
	if err := mysql.Init(); err != nil {
		fmt.Printf("init mysql failde, err:%v\n", err)
	}
	defer mysql.Close()

	//初始化redis连接
	if err := redis.Init(); err != nil {
		fmt.Printf("init redix failde, err:%v\n", err)
	}
	defer redis.Close()

	//初始化ID生成器

	//注册路由
	r := route.Setup()

	//启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf("%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func(){
		//开启一个goroutine启动服务
		if err:=srv.ListenAndServe();err!=nil&&err!=http.ErrServerClosed{
			
		}
	}
}
