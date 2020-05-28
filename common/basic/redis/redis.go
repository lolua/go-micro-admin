package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	log "github.com/micro/go-micro/v2/logger"
	"micro-admin/common/basic/config"
	"sync"
)

var (
	client *redis.Client
	m      sync.RWMutex
	inited bool
)

func init() {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Info("已经初始化过Redis...")
		return
	}

	redisConfig := config.GetRedisConfig()

	// 打开才加载
	if redisConfig != nil && redisConfig.GetEnabled() {
		log.Info("初始化Redis...")

		// 加载哨兵模式
		log.Info("初始化Redis，普通模式...")
		initSingle(redisConfig)

		log.Info("初始化Redis，检测连接...")

		pong, err := client.Ping(context.TODO()).Result()
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Info("初始化Redis，检测连接Ping.")
		log.Info("初始化Redis，检测连接Ping..")
		log.Infof("初始化Redis，检测连接Ping... %s", pong)
	}
	inited = true
}

// GetRedis 获取redis
func GetRedis() *redis.Client {
	return client
}

func initSingle(redisConfig config.RedisConfig) {
	client = redis.NewClient(&redis.Options{
		Addr:     redisConfig.GetConn(),
		Password: redisConfig.GetPassword(), // no password set
		DB:       redisConfig.GetDBNum(),    // use default DB
	})
}
