//Package log 日志模块.
//	向控制台和文件输出日志信息
//	[DEBUG]  调试类型日志只向控制台打印
//	[SIGNAL] 信号类型日志，打印级别是最高的
package log

import (
	"github.com/longbozhan/timewriter"
	"kada/server/utils/config"
	"log"
	"time"
)

type Level int

const (
	_ Level = iota
	LSignal
	LError
	LWarn
	LInfo
	LDebug
	LCrash
)

var (
	_level  Level // 日志等级
	_logger *log.Logger
)

// 启动日志
func Start() {
	_level = LDebug
	if level, err := config.ToInt(config.Get(config.Logger, config.LoggerLevel)); err == nil {
		_level = Level(level)
	}
	
	writer := &timewriter.TimeWriter{
		Dir:        "./logs",
		Compress:   true,
		ReserveDay: 30,
	}
	_logger = log.New(writer, "", 0)
}

// 写入日志
func Write(level string, s string, v ...interface{}) {
	ts := time.Now().Format("15:04:05.999999") //设定时间格式
	var a []interface{}
	a = append(a, ts)
	a = append(a, s)
	a = append(a, v...)
	_logger.SetPrefix(level)
	_logger.Printf(" %-15s %s", a...)
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
