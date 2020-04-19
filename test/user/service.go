package user

import "kada/server/core/service"

var (
	_service *service.Service
)

func init() {
	_service = service.NewService()
	_service.Register("user", new(Handler))
	_service.Start()
}

// 调用服务
func Call(action string, args interface{}, back interface{}) error {
	return _service.Call("user", action, args, back)
}
