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

func output(args interface{}) error {
	if err := _service.Call(MODULE, "Log", args, nil); err != nil {
		return err
	}
	return nil
}

func Load(filename string) error {
	args := &LoadArgs{filename}
	if err := _service.Call(MODULE, "Load", args, nil); err != nil {
		return err
	}
	return nil
}

// 打印调试级别日志
func Debug(format string, a ...interface{}) error {
	return output(&LogArgs{LOG_DEBUG, fmt.Sprintf(format, a...)})
}

// 打印信息级别日志
func Info(format string, a ...interface{}) error {
	return output(&LogArgs{LOG_INFO, fmt.Sprintf(format, a...)})
}

// 打印警告级别日志
func Warn(format string, a ...interface{}) error {
	return output(&LogArgs{LOG_WARN, fmt.Sprintf(format, a...)})
}

// 打印错误级别日志
func Error(format string, a ...interface{}) error {
	return output(&LogArgs{LOG_ERROR, fmt.Sprintf(format, a...)})
}

// 打印信号级别日志
func Signal(format string, a ...interface{}) error {
	return output(&LogArgs{LOG_SIGNAL, fmt.Sprintf(format, a...)})
}

// 打印崩溃级别日志（格式化）
func Dump(format string, a ...interface{}) error {
	return output(&LogArgs{LOG_CRASH, fmt.Sprintf(format, a...)})
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
