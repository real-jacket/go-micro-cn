package main

import (
	"fmt"
	"time"

	"github.com/go-micro-cn/tutorials/basic"
	"github.com/go-micro-cn/tutorials/basic/config"
	"github.com/go-micro-cn/tutorials/user-service/handler"
	"github.com/go-micro-cn/tutorials/user-service/model"
	s "github.com/go-micro-cn/tutorials/user-service/proto/user"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	// New Service
	service := micro.NewService(
		micro.Name("mu.micro.book.service.user"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) error {
			// 初始化模型层
			model.Init()
			// 初始化handler
			handler.Init()
			return nil
		}),
	)

	// 注册服务
	//这里如果使用的是v1的protoc-gen-micro的话，会编译报错。使用`go get -u github.com/micro/protoc-gen-micro/v2`安装v2的新插件即可。
	s.RegisterUserHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Timeout = time.Second * 5
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
