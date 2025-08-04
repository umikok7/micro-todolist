package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"task/conf"
	"task/core"
	"task/services"
)

func main() {
	conf.Init()
	// etcd 注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs("192.168.1.107:12379"),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcTaskService"), // 微服务名字
		micro.Address("127.0.0.1:8083"),
		micro.Registry(etcdReg),
	)

	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = services.RegisterTaskServiceHandler(microService.Server(), new(core.TaskService))
	_ = microService.Run()
}
