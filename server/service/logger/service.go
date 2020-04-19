package logger

import (
	"fmt"
	"kada/server/core/service"
)

const (
	MODULE = "logger"
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

// 打印崩溃级别日志
func Panic(format string, a ...interface{}) error {
	return output(&LogArgs{LOG_CRASH, fmt.Sprintf(format, a...)})
}
