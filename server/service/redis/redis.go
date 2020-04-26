package redis

import (
	"fmt"
	"kada/server/log"
	"time"

	"github.com/gomodule/redigo/redis"
)

const (
	Min = "-inf"
	Max = "+inf"
)

var (
	String    = redis.String
	Strings   = redis.Strings
	StringMap = redis.StringMap
	Int       = redis.Int
	Ints      = redis.Ints
	IntMap    = redis.IntMap
	Int64     = redis.Int64
	Int64s    = redis.Int64s
	Int64Map  = redis.Int64Map
	Bool      = redis.Bool
)

type Client struct {
	Pool *redis.Pool
}

func NewClient(host string, port int, db int, password string) *Client {
	client := new(Client)
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
				redis.DialDatabase(db),
				redis.DialPassword(password),
			)
		},
	}
	return client
}

func (o *Client) Exec(cmd string, key interface{}, args ...interface{}) (interface{}, error) {
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

func (o *Client) Get(key string, def string) string {
	ret, err := String(o.Exec("GET", key))
	if err != nil {
		log.Printf("[redis] get %v err: %v", key, err)
		return def
	}
	return ret
}

func (o *Client) HGet(key string, field string, def string) string {
	ret, err := String(o.Exec("HGET", key, field))
	if err != nil {
		log.Printf("[redis] hget %s %s err: %v", key, field, err)
		return def
	}
	return ret
}

func (o *Client) HGetAll(key string) map[string]string {
	ret, err := StringMap(o.Exec("HGETALL", key))
	if err != nil {
		log.Printf("[redis] hgetall %s err: %v", key, err)
		return nil
	}
	return ret
}

func (o *Client) ZRANGEBYSCORE(zset string, min string, max string) map[string]string {
	ret, err := StringMap(o.Exec("ZRANGEBYSCORE", zset, min, max))
	if err != nil {
		log.Warn("[redis] zrangebyscore %s %s %s err: %v", zset, min, max, err)
		return nil
	}
	return ret
}

func (o *Client) NoInclude(num int) string {
	return fmt.Sprintf("(%d", num)
}
