package log

import (
	"fmt"
	"kada/server/core/service"
	"log"
)

const (
	MODULE = "log"
)

var (
	_service *service.Service
)

func init() {
	_service = service.NewService()
	_service.Register(MODULE, new(Handler))
	_service.Start()
}

func Call(args *Args) error {
	return _service.Call(MODULE, "Log", args, nil)
}

// 打印调试级别日志
func Debug(format string, a ...interface{}) error {
	return Call(&Args{LvDebug, fmt.Sprintf(format, a...)})
}

// 打印信息级别日志
func Info(format string, a ...interface{}) error {
	return Call(&Args{LvInfo, fmt.Sprintf(format, a...)})
}

// 打印警告级别日志
func Warn(format string, a ...interface{}) error {
	return Call(&Args{LvWarn, fmt.Sprintf(format, a...)})
}

// 打印错误级别日志
func Error(format string, a ...interface{}) error {
	return Call(&Args{LvError, fmt.Sprintf(format, a...)})
}

// 打印信号级别日志
func Signal(format string, a ...interface{}) error {
	return Call(&Args{LvSignal, fmt.Sprintf(format, a...)})
}

// 打印崩溃级别日志（格式化）
func Dump(format string, a ...interface{}) error {
	return Call(&Args{LvCrash, fmt.Sprintf(format, a...)})
}

// 兼容日志崩溃函数
func Panic(a ...interface{}) {
	log.Panic(a...)
}

// 兼容日志打印函数
func Print(a ...interface{}) {
	log.Print(a...)
}

// 兼容日志打印函数（格式化）
func Printf(format string, a ...interface{}) {
	log.Printf(format, a...)
}

// 兼容日志打印函数
func Fatal(a ...interface{}) {
	log.Fatal(a...)
}
