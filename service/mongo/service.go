package mongo

import (
	"fmt"
)

const (
	NoFoundError string = "mongo: no documents in result"
)

var (
	dbs map[string]*MongoDB
)

func init() {
	dbs = make(map[string]*MongoDB)
}

// 启动服务
func Start(uri string, db string) error {
	mongo := new(MongoDB)
	if err := mongo.Connect(uri, db); err != nil {
		return fmt.Errorf("[mongo] connect uri(%s) db(%s) err: %w", uri, db, err)
	}
	dbs[db] = mongo
	return nil
}

func Db(name string) *MongoDB {
	return dbs[name]
}
