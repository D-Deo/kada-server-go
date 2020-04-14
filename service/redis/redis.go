package redis

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	Pool *redis.Pool
}

func (o *Redis) Exec(cmd string, key interface{}, args ...interface{}) (interface{}, error) {
	conn := o.Pool.Get()
	if err := conn.Err(); err != nil {
		return nil, fmt.Errorf("[redis] conn get err: %w", err)
	}
	defer conn.Close()
	params := make([]interface{}, 0)
	params = append(params, key)

	if len(args) > 0 {
		for _, v := range args {
			params = append(params, v)
		}
	}
	reply, err := conn.Do(cmd, params...)
	if err != nil {
		return nil, fmt.Errorf("[redis] conn do err: %w", err)
	}
	return reply, nil
}

func (o *Redis) HGetAll(key string) map[string]string {
	ret, err := redis.StringMap(o.Exec("HGETALL", key))
	if err != nil {
		log.Printf("[redis] hgetall %s err: %v", key, err)
		return nil
	}
	return ret
}
