package log

import (
	"fmt"
	"kada/server/core"
)

const (
	MODULE = "log"
)

var (
	_service *core.Service
)

func init() {
	_service = core.NewService()
	_service.Register(MODULE, new(Handler))
	_service.Start()
}

func Call(args *Args) error {
	return _service.Call(MODULE, "Log", args, nil)
}

// 打印调试级别日志
func Debug(format string, a ...interface{}) error {
	return Call(&Args{LDebug, fmt.Sprintf(format, a...)})
}

// 打印信息级别日志
func Info(format string, a ...interface{}) error {
	return Call(&Args{LInfo, fmt.Sprintf(format, a...)})
}

// 打印警告级别日志
func Warn(format string, a ...interface{}) error {
	return Call(&Args{LWarn, fmt.Sprintf(format, a...)})
}

// 打印错误级别日志
func Error(format string, a ...interface{}) error {
	return Call(&Args{LError, fmt.Sprintf(format, a...)})
}

// 打印信号级别日志
func Signal(format string, a ...interface{}) error {
	return Call(&Args{LSignal, fmt.Sprintf(format, a...)})
}

// 打印崩溃级别日志（格式化）
func Dump(format string, a ...interface{}) error {
	return Call(&Args{LCrash, fmt.Sprintf(format, a...)})
}
