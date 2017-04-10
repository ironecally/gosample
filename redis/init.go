package redis

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var RedisPool redisPool

type redisPool struct {
	redis1 *redis.Pool
}

func InitRedis() {
	RedisPool = redisPool{
		redis1: getRedisConn(),
	}
}

func getRedisConn() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 10 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", "devel-redis.tkpd:6379") },
	}
}
