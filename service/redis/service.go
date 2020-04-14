package redis

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	clients map[string]*Redis
)

func init() {
	clients = make(map[string]*Redis)
}

func NewRedis(host string, port int32, db int32, password string) *Redis {
	client := new(Redis)
	client.Pool = &redis.Pool{
		MaxIdle:     256,
		MaxActive:   0,
		IdleTimeout: time.Duration(120),
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				fmt.Sprintf("%s:%d", host, port),
				redis.DialReadTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialDatabase(int(db)),
				redis.DialPassword(password),
			)
		},
	}
	return client
}

func Set(name string, client *Redis) {
	clients[name] = client
}

func Get(name string) *Redis {
	return clients[name]
}
