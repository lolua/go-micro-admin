package main

import (
	"fmt"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"micro-admin/common/basic/config"
	"micro-admin/user/user-web/handler"
	"net/http"
)

func main() {

	reg := etcd.NewRegistry(func(op *registry.Options) {
		etcdConfig := config.GetEtcdConfig()
		op.Addrs = []string{
			fmt.Sprintf("%s:%d", etcdConfig.GetHost(), etcdConfig.GetPort()),
		}
	})
	// create new web service
	service := web.NewService(
		web.Name("zhj.micro.admin.web.user"),
		web.Registry(reg),
		web.Version("latest"),
		web.Address(":8089"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/user/call", handler.Login)
	service.HandleFunc("/user/get", handler.Get)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
