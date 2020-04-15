package redis

import "log"

var (
	clients map[string]*Client
)

func init() {
	clients = make(map[string]*Client)
}

func Set(name string, client *Client) {
	clients[name] = client
}

func Get(name string) *Client {
	client, ok := clients[name]
	if !ok {
		log.Panicf("[redis] no found redis client: %s", name)
		return nil
	}
	return client
}
