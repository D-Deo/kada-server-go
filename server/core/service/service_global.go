package service

var (
	_service *Service
)

func init() {
	_service = NewService()
}

// 注册控制器
func Register(name string, handler interface{}) {
	_service.Register(name, handler)
}

// 启动服务
func Start() {
	_service.Start()
}

// 调用服务
func Call(handle string, action string, args interface{}, back interface{}) error {
	return _service.Call(handle, action, args, back)
}
