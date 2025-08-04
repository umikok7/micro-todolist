package main

import (
	"mq-server/conf"
	"mq-server/service"
)

func main() {
	conf.Init() // 数据库和MQ初始化

	forever := make(chan bool)
	service.CreateTask()
	<-forever // 往forever中读数据，但是没有如何地方写入，所以此时主 goroutine 在这里阻塞
}
