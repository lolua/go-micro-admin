package access

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	r "github.com/go-redis/redis/v8"
	"micro-admin/auth/auth-srv/model/admin_user"
	jwtUtil "micro-admin/auth/auth-srv/model/jwt"
	"micro-admin/common/basic/redis"
	"strconv"
	"sync"
	"time"
)

// service 服务
type service struct {
}

var (
	s                *service
	redisClient      *r.Client
	m                sync.RWMutex
	authTokenPrefix  = "auth.token.id:"
	tokenExpiredDate = 3600 * 24 * 30 * time.Second
)

// MakeAccessToken 生成token
func (s *service) CreateAccessToken(a *admin_user.AdminUser) (ret string, err error) {
	fmt.Println(a)
	userId := strconv.FormatInt(a.Id, 10)
	j, e := json.Marshal(a)
	if e != nil {
		return "", errors.New("生成token失败")
	}
	tk, err := jwtUtil.GenerateToken(&jwtUtil.Subject{ID: userId, Data: string(j)})
	if err != nil {
		return "", err
	}
	a.Token = tk
	j, e = json.Marshal(a)
	if e != nil {
		return "", errors.New("生成token失败")
	}
	redisClient.Set(context.TODO(), authTokenPrefix+userId, string(j), tokenExpiredDate)
	return tk, nil
}

// GetCachedAccessToken 获取缓存的token
func (s *service) ValidAccessToken(tk string) (err error) {
	subject, err := jwtUtil.GetSubjectFromToken(tk)
	if err != nil {
		return err
	}
	stcmd := redisClient.Get(context.TODO(), authTokenPrefix+subject.ID)
	m, err := stcmd.Result()
	if stcmd.Err() == nil && err == nil {
		var a admin_user.AdminUser
		if json.Unmarshal([]byte(m), &a) == nil {
			if a.Token != tk {
				return errors.New("账号异地登录")
			}
			return nil
		}
	} else if err != nil {
		return errors.New("长时间未登录")
	}
	return errors.New("账户验证失败,err: " + stcmd.Err().Error())
}

// DelUserAccessToken 清除用户token
func (s *service) DelAccessToken(token string) (err error) {
	subject, err := jwtUtil.GetSubjectFromToken(token)
	if err != nil {
		return
	}
	redisClient.Del(context.TODO(), authTokenPrefix+subject.ID)
	return
}

func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

func init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	redisClient = redis.GetRedis()

	s = &service{}
}
