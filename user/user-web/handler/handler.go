package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/v2/client"
	user "micro-admin/user/user-srv/proto/user"
	"net/http"
)

var (
	userClient = user.NewUserService("zhj.micro.admin.srv.user", client.DefaultClient)
)

func Login(w http.ResponseWriter, r *http.Request) {

	rsp, err := userClient.Login(context.TODO(), &user.Request{
		Name: "myna1",
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println(rsp)
	ret, e := json.Marshal(rsp)
	if e != nil {
		http.Error(w, e.Error(), 500)
	} else {
		w.Write(ret)
	}
}
func Get(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	rsp, err := userClient.Get(context.TODO(), &user.Request{
		Name: values.Get("token"),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println(rsp)
	ret, e := json.Marshal(rsp)
	if e != nil {
		http.Error(w, e.Error(), 500)
	} else {
		w.Write(ret)
	}
}
