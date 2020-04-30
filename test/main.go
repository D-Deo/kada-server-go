package main

import (
	"kada/server/console"
	"kada/server/core"
	"kada/server/log"
	"kada/server/utils/config"
	"kada/test/user"
	"runtime"
)

func show(args ...string) {
	log.Warn("[main] server status NumCPU(%d) NumGoroutine(%d) NumCgoCall(%d)", runtime.NumCPU(), runtime.NumGoroutine(), runtime.NumCgoCall())
}

func test1(args ...string) {
	req := &user.CreateArgs{}
	req.UserName = "test"
	req.Password = "123123"
	req.IP = "127.0.0.1"
	req.Device = 1
	req.Phone = "12345678901"
	req.Token = "token_1"

	var ret user.CreateBack
	if err := user.Call("Create", req, &ret); err != nil {
		log.Panic("[test2] user create err: %v", err)
	}

	log.Info("[test2] user create ret: %v", ret.User)
}

func main() {
	defer core.Panic()

	// 利用cpu多核来控制go的协程
	ncpu := runtime.NumCPU()
	ngoc := runtime.GOMAXPROCS(ncpu)

	// 初始化配置
	if err := config.Load("test"); err != nil {
		log.Panic(err)
	}

	log.Signal("[main] monitor startup CPU: %d GOMAXPROC: %d", ncpu, ngoc)

	console.Register("show", show)
	console.Register("test1", test1)
	console.Listen()
}
