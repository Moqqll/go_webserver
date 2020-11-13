package main

import (
	"fmt"
	"go_webserver/logger"
	"go_webserver/setting"
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
	//初始化mysql连接
	if err := mysql.Init(); err != nil {
		fmt.Printf("init mysql failde, err:%v\n", err)
	}
	//初始化redis连接
	if err := redix.Init(); err != nil {
		fmt.Printf("init redix failde, err:%v\n", err)
	}
	//初始化ID生成器

	//注册路由

	//启动服务（优雅关机）
}
