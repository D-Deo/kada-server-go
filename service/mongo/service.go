package mongo

import (
	"fmt"
	"log"
)

const (
	NoFoundError string = "mongo: no documents in result"
)

var (
	clients map[string]*Client
)

func init() {
	clients = make(map[string]*Client)
}

// 连接
func Connect(uri string, db string) error {
	mongo := new(Client)
	if err := mongo.Connect(uri, db); err != nil {
		return fmt.Errorf("[mongo] connect uri(%s) db(%s) err: %w", uri, db, err)
	}
	clients[db] = mongo
	return nil
}

// 获取
func Get(name string) *Client {
	client, ok := clients[name]
	if !ok {
		log.Panicf("[redis] no found redis client: %s", name)
		return nil
	}
	return client
}
