package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"micro-admin/common/basic/config"
	"micro-admin/user/user-srv/handler"

	user "micro-admin/user/user-srv/proto/user"
)

func main() {

	reg := etcd.NewRegistry(func(op *registry.Options) {
		etcdConfig := config.GetEtcdConfig()
		op.Addrs = []string{
			fmt.Sprintf("%s:%d", etcdConfig.GetHost(), etcdConfig.GetPort()),
		}
	})

	// New Service
	service := micro.NewService(
		micro.Name("zhj.micro.admin.srv.user"),
		micro.Registry(reg),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
