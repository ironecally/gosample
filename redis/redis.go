package redis

import (
	"fmt"
	"net/http"

	"github.com/garyburd/redigo/redis"
)

func SetRedisHandler(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	val := r.FormValue("value")

	redisConn := RedisPool.redis1.Get()
	redisConn.Do("SET", key, val)

	fmt.Fprintf(w, "OK")
}

func GetRedisHandler(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")

	redisConn := RedisPool.redis1.Get()
	res, err := redis.String(redisConn.Do("GET", key))
	if err != nil {
		fmt.Fprintf(w, "err %s", err.Error())
		return
	}

	fmt.Fprintf(w, "%s:%s", key, res)
	return
}
