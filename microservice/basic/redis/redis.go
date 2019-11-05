/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-18 08:20:02
 * @LastEditTime: 2019-08-18 08:27:52
 * @LastEditors: Please set LastEditors
 */
package redis

import (
	"sync"

	"github.com/go-redis/redis"
	"github.com/micro/go-micro/util/log"
	"github.com/yuwe1/micolearn/microservice/basic/config"
)

var (
	client *redis.Client
	m      sync.RWMutex
	inited bool
)

func Init() {
	m.Lock()
	defer m.Unlock()
	if inited {
		log.Log("已经初始化过Redis")
		return
	}
	redisConfig := config.GetRedisConfig()
	if redisConfig != nil && redisConfig.GetEnabled() {
		log.Log("初始化Redis")
		// 加载哨兵模式
		if redisConfig.GetSentinelConfig() != nil && redisConfig.GetSentinelConfig().GetEnabled() {
			log.Log("哨兵模式")
			initSentinel(redisConfig)
		} else { //普通模式
			log.Log("初始化redis，普通模式")
			initSingle(redisConfig)
		}
		log.Log("初始化redis，检测连接")
		pong, err := client.Ping().Result()
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Log("初始化Redis，检测连接Ping.")
		log.Log("初始化Redis，检测连接Ping..")
		log.Logf("初始化Redis，检测连接Ping... %s", pong)
	}
}

// GetRedis 获取redis
func GetRedis() *redis.Client {
	return client
}

func initSentinel(redisConfig config.RedisConfig) {
	client = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    redisConfig.GetSentinelConfig().GetMaster(),
		SentinelAddrs: redisConfig.GetSentinelConfig().GetNodes(),
		DB:            redisConfig.GetDBNum(),
		Password:      redisConfig.GetPassword(),
	})

}

func initSingle(redisConfig config.RedisConfig) {
	client = redis.NewClient(&redis.Options{
		Addr:     redisConfig.GetConn(),
		Password: redisConfig.GetPassword(), // no password set
		DB:       redisConfig.GetDBNum(),    // use default DB
	})
}
