package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/common/log"
	"micro-admin/common/basic/db"
	rds "micro-admin/common/basic/redis"
	"time"
)

type User struct {
	Id         int64     `form:"id" gorm:"primary_key"`
	Name       string    `form:"name" gorm:"index;unique;not null;type:varchar(100)"`
	Card       int64     `form:"card" gorm:"not null"`
	Caa        int64     `form:"caa" gorm:"not null;column:cards;default:0"`
	Ces        int64     `form:"caa" gorm:"not null;column:cards1;default:0"`
	CreateTime time.Time `form:"createTime" gorm:"not null" time_format:"2006-01-02" time_utc:"1"`
}

var (
	userDB      *gorm.DB
	redisClient *redis.Client
)

func (u *User) Insert() (int64, error) {
	ret := userDB.Create(u)
	if len(ret.GetErrors()) > 0 {
		log.Error(ret.GetErrors())
		return 0, errors.New(fmt.Sprint(ret.GetErrors()))
	}
	return u.Id, nil
}
func (u *User) Update() (int64, error) {
	ret := userDB.Save(u)
	if len(ret.GetErrors()) > 0 {
		log.Error(ret.GetErrors())
		return 0, errors.New(fmt.Sprint(ret.GetErrors()))
	}
	return u.Id, nil
}

func (u *User) FindOne() (*User, error) {
	m := redisClient.Get(context.TODO(), "user:name:"+u.Name)
	fmt.Println(m)
	if m.Err() == nil {
		s, e := m.Result()
		if e == nil {
			var uu User
			e := json.Unmarshal([]byte(s), &uu)
			if e == nil {
				return &uu, nil
			} else {
				log.Error(e)
			}
		} else {
			log.Error(e)
		}
	} else {
		log.Error(m.Err())
	}
	ret := userDB.Where(u).First(u)
	if ret.RecordNotFound() {
		return nil, gorm.ErrRecordNotFound
	}
	if len(ret.GetErrors()) > 0 {
		log.Error(ret.GetErrors())
		return nil, errors.New(fmt.Sprint(ret.GetErrors()))
	}
	t, e := json.Marshal(u)
	if e == nil {
		redisClient.Set(context.TODO(), "user:name:"+u.Name, t, time.Hour)
	} else {
		log.Error(e)
	}
	fmt.Println(u)
	return u, nil
}

func init() {
	table := db.LoginDB().HasTable(User{})
	if !table {
		defer func() {
			if err := recover(); err != nil {
				panic(err)
			}
		}()
		d := db.LoginDB().CreateTable(User{})
		if len(d.GetErrors()) > 0 {
			panic(fmt.Sprint(d.GetErrors()))
		}
	}
	userDB = db.LoginDB()
	redisClient = rds.GetRedis()
}
