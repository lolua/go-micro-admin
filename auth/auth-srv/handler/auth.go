package handler

import (
	"context"
	"micro-admin/auth/auth-srv/model/access"
	"micro-admin/auth/auth-srv/model/admin_user"

	log "github.com/micro/go-micro/v2/logger"

	auth "micro-admin/auth/auth-srv/proto/auth"
)

var accessService access.Service

type Auth struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Auth) CreateAccessToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Info("Received Auth.CreateAccessToken request")
	ret, err := accessService.CreateAccessToken(&admin_user.AdminUser{Id: req.Id, Name: req.Name})
	if err != nil {
		rsp.Code = 400
		rsp.Msg = err.Error()
	} else {
		rsp.Code = 200
		rsp.Msg = "生成token成功"
		rsp.Token = ret
	}
	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Auth) DelAccessToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Info("Received Auth.DelAccessToken request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Auth) RefreshAccessToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Info("Received Auth.RefreshAccessToken request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Auth) ValidAccessToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Info("Received Auth.ValidAccessToken request")
	err := accessService.ValidAccessToken(req.Token)
	if err != nil {
		rsp.Code = 400
		rsp.Msg = err.Error()
	} else {
		rsp.Code = 200
		rsp.Msg = "SUccess"
	}
	return nil
}

func init() {
	var err error
	accessService, err = access.GetService()
	if err != nil {
		log.Fatal(err)
	}
}
