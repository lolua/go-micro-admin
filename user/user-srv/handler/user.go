package handler

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/client"
	user "micro-admin/user/user-srv/proto/user"

	log "github.com/micro/go-micro/v2/logger"

	auth "micro-admin/auth/auth-srv/proto/auth"
	us "micro-admin/user/user-srv/model/user"
)

type User struct{}

var (
	authClient = auth.NewAuthService("zhj.micro.admin.srv.auth", client.DefaultClient)
)

// Call is a single request handler called via client.Call or the generated client code
func (e *User) Get(ctx context.Context, req *user.Request, rsp *user.Response) error {
	response, err := authClient.ValidAccessToken(context.TODO(), &auth.Request{Token: req.Name})
	if err != nil {
		rsp.Msg = err.Error()
		rsp.Code = 400
	} else if response.Code != 200 {
		rsp.Code = 400
		rsp.Msg = response.Msg
	} else {
		rsp.Msg = "查询成功"
		rsp.Code = 200
	}
	return nil
}
func (e *User) Login(ctx context.Context, req *user.Request, rsp *user.Response) error {
	log.Info("Received User.Call request")
	u := us.User{Name: req.Name}
	ret, err := u.FindOne()
	if err != nil {
		rsp.Msg = err.Error()
		rsp.Code = 400
	} else {
		response, err := authClient.CreateAccessToken(context.TODO(), &auth.Request{Id: ret.Id, Name: ret.Name})
		if err != nil {
			rsp.Msg = err.Error()
			rsp.Code = 400
		} else if response.Code != 200 {
			rsp.Code = 400
			rsp.Msg = response.Msg
		} else {
			rsp.Msg = "查询成功"
			rsp.Code = 200
			rsp.Data = &user.Message{
				Id:         ret.Id,
				Name:       ret.Name,
				CreateTime: ret.CreateTime.Format("2006-01-02 15:04:05"),
				Card:       ret.Card,
				Token:      response.Token,
			}
		}
	}
	fmt.Println(rsp)
	return nil
}
