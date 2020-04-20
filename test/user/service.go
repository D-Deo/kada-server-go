package user

import (
	"kada/server/core"
)

var (
	_service *core.Service
)

func init() {
	_service = core.NewService()
	_service.Register("user", new(Handler))
	_service.Start()
}

// 调用服务
func Call(action string, args interface{}, back interface{}) error {
	return _service.Call("user", action, args, back)
}
