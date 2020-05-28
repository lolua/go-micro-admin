package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"micro-admin/common/basic/config"
	"sync"

	log "github.com/micro/go-micro/v2/logger"
)

var (
	inited  bool
	loginDB *gorm.DB
	gmDB    *gorm.DB
	hallDB  *gorm.DB
	m       sync.RWMutex
)

// Init 初始化数据库
func init() {
	m.Lock()
	defer m.Unlock()

	var err error

	if inited {
		err = fmt.Errorf("[Init] db 已经初始化过")
		log.Error(err)
		return
	}

	if config.GetLoginMysqlConfig().GetEnabled() {
		loginDB = initMysql(config.GetLoginMysqlConfig())
	}
	if config.GetGmMysqlConfig().GetEnabled() {
		gmDB = initMysql(config.GetGmMysqlConfig())
	}
	if config.GetHallMysqlConfig().GetEnabled() {
		hallDB = initMysql(config.GetHallMysqlConfig())
	}

	inited = true
}

// GetDB 获取db
func LoginDB() *gorm.DB {
	return loginDB
}
func GmDB() *gorm.DB {
	return gmDB
}
func HallDB() *gorm.DB {
	return hallDB
}
