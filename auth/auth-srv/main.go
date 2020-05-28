package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"micro-admin/auth/auth-srv/handler"
	"micro-admin/common/basic/config"

	auth "micro-admin/auth/auth-srv/proto/auth"
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
		micro.Name("zhj.micro.admin.srv.auth"),
		micro.Registry(reg),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	auth.RegisterAuthHandler(service.Server(), new(handler.Auth))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
