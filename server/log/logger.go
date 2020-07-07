package log

import (
	"encoding/json"
	"fmt"
	"github.com/longbozhan/timewriter"
	"io"
	"log"
	"time"
)

type Fields map[string]interface{}

type Option struct {
	Level   Level
	Console bool
}

type Logger struct {
	Fields *Fields
	Option Option
	Logger *log.Logger
}

func NewLogger(level Level, console bool) *Logger {
	logger := &Logger{}
	logger.Option.Level = level
	logger.Option.Console = console
	writer := &timewriter.TimeWriter{
		Dir:        "./logs",
		Compress:   true,
		ReserveDay: 30,
	}
	logger.Logger = log.New(writer, "", 0)
	return logger
}

func (o *Logger) Writer() io.Writer {
	return o.Logger.Writer()
}

// 结构化
func (o *Logger) With(fields Fields) *Logger {
	o.Fields = &fields
	return o
}

func (o *Logger) checkFields(format string, v ...interface{}) (string, interface{}) {
	var fields []byte
	if o.Fields != nil {
		fields, _ = json.Marshal(o.Fields)
	}
	if len(fields) > 0 {
		format += " %s"
		v = append(v, string(fields))
	}
	o.Fields = nil
	return format, v
}

// 输出日志
func (o *Logger) Output(level Level, output string) {
	ts := time.Now().Format("2006-01-02 15:04:05.999999") //设定时间格式

	switch level {
	case LDebug:
		o.Write(level, "[D]", output, ts)
	case LInfo:
		o.Write(level, "[I]", output, ts)
	case LSignal:
		o.Write(level, "[S]", output, ts)
	case LWarn:
		o.Write(level, "[W]", output, ts)
	case LError:
		o.Write(level, "[E]", output, ts)
	default:
		break
	}
}

// 写入日志
func (o *Logger) Write(level Level, tag string, output string, ts string) {
	if o.Option.Level >= level {
		if o.Option.Console {
			fmt.Printf("%s %-26s %s \n", tag, ts, output)
		}
		var a []interface{}
		a = append(a, ts)
		a = append(a, output)
		o.Logger.SetPrefix(tag)
		o.Logger.Printf(" %-15s %s", a...)
	}
}

// 打印调试级别日志
func (o *Logger) Debug(format string, v ...interface{}) {
	var fields []byte
	if o.Fields != nil {
		fields, _ = json.Marshal(o.Fields)
	}
	if len(fields) > 0 {
		format += " %s"
		v = append(v, string(fields))
	}
	o.Fields = nil
	go o.Output(LDebug, fmt.Sprintf(format, v...))
}

// 打印信息级别日志
func (o *Logger) Info(format string, v ...interface{}) {
	var fields []byte
	if o.Fields != nil {
		fields, _ = json.Marshal(o.Fields)
	}
	if len(fields) > 0 {
		format += " %s"
		v = append(v, string(fields))
	}
	o.Fields = nil
	go o.Output(LInfo, fmt.Sprintf(format, v...))
}

// 打印警告级别日志
func (o *Logger) Warn(format string, v ...interface{}) {
	var fields []byte
	if o.Fields != nil {
		fields, _ = json.Marshal(o.Fields)
	}
	if len(fields) > 0 {
		format += " %s"
		v = append(v, string(fields))
	}
	o.Fields = nil
	go o.Output(LWarn, fmt.Sprintf(format, v...))
}

// 打印错误级别日志
func (o *Logger) Error(format string, v ...interface{}) {
	var fields []byte
	if o.Fields != nil {
		fields, _ = json.Marshal(o.Fields)
	}
	if len(fields) > 0 {
		format += " %s"
		v = append(v, string(fields))
	}
	o.Fields = nil
	go o.Output(LError, fmt.Sprintf(format, v...))
}

// 打印信号级别日志
func (o *Logger) Signal(format string, v ...interface{}) {
	var fields []byte
	if o.Fields != nil {
		fields, _ = json.Marshal(o.Fields)
	}
	if len(fields) > 0 {
		format += " %s"
		v = append(v, string(fields))
	}
	o.Fields = nil
	go o.Output(LSignal, fmt.Sprintf(format, v...))
}

// 兼容日志打印函数
func (o *Logger) Print(v ...interface{}) {
	var fields []byte
	if o.Fields != nil {
		fields, _ = json.Marshal(o.Fields)
	}
	if len(fields) > 0 {
		v = append(v, string(fields))
	}
	o.Fields = nil
	go o.Output(LDebug, fmt.Sprint(v...))
}

// 兼容日志打印函数（格式化）
func (o *Logger) Printf(format string, v ...interface{}) {
	o.Debug(format, v...)
}
