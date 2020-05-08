package log

import (
	"log"
)

type Level int

const (
	_ Level = iota
	LSignal
	LError
	LWarn
	LInfo
	LDebug
)

var (
	_logger *Logger
)

func init() {
	_logger = NewLogger(LDebug, true)
}

// 设置参数
func SetOptions(level Level, console bool) {
	_logger.Option.Level = level
	_logger.Option.Console = console
}

// 结构化
func With(fields Fields) *Logger {
	return _logger.With(fields)
}

// 打印调试级别日志
func Debug(format string, v ...interface{}) {
	_logger.Debug(format, v...)
}

// 打印信息级别日志
func Info(format string, v ...interface{}) {
	_logger.Info(format, v...)
}

// 打印警告级别日志
func Warn(format string, v ...interface{}) {
	_logger.Warn(format, v...)
}

// 打印错误级别日志
func Error(format string, v ...interface{}) {
	_logger.Error(format, v...)
}

// 打印信号级别日志
func Signal(format string, v ...interface{}) {
	_logger.Signal(format, v...)
}

// 兼容日志打印函数
func Print(v ...interface{}) {
	_logger.Print(v...)
}

// 兼容日志打印函数（格式化）
func Printf(format string, v ...interface{}) {
	_logger.Debug(format, v...)
}

// 兼容日志崩溃函数
func Panic(v ...interface{}) {
	log.Panic(v...)
}

// 兼容日志打印函数
func Fatal(v ...interface{}) {
	log.Fatal(v...)
}
