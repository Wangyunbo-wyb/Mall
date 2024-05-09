package main

import (
	conf "Mall/config"
	util "Mall/pkg/utils/log"
	"Mall/pkg/utils/track"
	"Mall/repository/cache"
	"Mall/repository/db/dao"
	"Mall/repository/es"
	"Mall/routes"
	"fmt"
)

func main() {
	loading() // 加载配置
	r := routes.NewRouter()
	_ = r.Run(conf.Config.System.HttpPort)
	fmt.Println("启动配成功...")
}

// loading一些配置
func loading() {
	conf.InitConfig()
	dao.InitMySQL()
	cache.InitCache()
	//rabbitmq.InitRabbitMQ() // 如果需要接入RabbitMQ可以打开这个注释
	es.InitEs() // 如果需要接入ELK可以打开这个注释
	//kafka.InitKafka()
	track.InitJaeger()
	util.InitLog() // 如果接入ELK请进入这个func打开注释
	fmt.Println("加载配置完成...")
	go scriptStarting()
}

func scriptStarting() {
	// 启动一些脚本
}
