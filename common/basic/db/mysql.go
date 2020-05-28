package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/micro/go-micro/v2/logger"
	"micro-admin/common/basic/config"
	"time"
)

func initMysql(mysqlConfig config.MysqlConfig) *gorm.DB {
	var err error

	// 创建连接
	mysqlDB, err := gorm.Open("mysql", mysqlConfig.GetURL())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	mysqlDB.SingularTable(true)
	mysqlDB.LogMode(true)
	// 最大连接数
	mysqlDB.DB().SetMaxOpenConns(mysqlConfig.GetMaxOpenConnection())
	// 最大闲置数
	mysqlDB.DB().SetMaxIdleConns(mysqlConfig.GetMaxIdleConnection())
	//连接数据库闲置断线的问题
	mysqlDB.DB().SetConnMaxLifetime(time.Second * mysqlConfig.GetConnMaxLifetime())
	// 激活链接
	if err = mysqlDB.DB().Ping(); err != nil {
		log.Fatal(err)
	}

	log.Infof("[DB] 启动Mysql： %s\n", mysqlConfig.GetURL())
	return mysqlDB
}
