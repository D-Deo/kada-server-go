//Package logger 日志模块.
//	向控制台和文件输出日志信息
//	[DEBUG]调试类型日志只向控制台打印
//	[SIGN]信号类型的日志，打印级别是最高的
package logger

import (
	"encoding/csv"
	"fmt"
	"log"
)

type LOG_LEVEL int

const (
	_ LOG_LEVEL = iota
	LOG_SIGNAL
	LOG_ERROR
	LOG_WARN
	LOG_INFO
	LOG_DEBUG
	LOG_CRASH
)

const (
	OUTPUT_LOG = "log"
	OUTPUT_CSV = "csv"
)

var (
	//LogLevel 控制台日志显示等级
	LogLevel LOG_LEVEL
	//LogOutput 输出类型
	LogOutput string
	
	//LogWriter 日志控制器
	LogWriter *log.Logger
	
	//CsvWriter 文件控制器
	CsvWriter *csv.Writer
)

//Init 初始化日志输出文件
func Init(filename string) error {
	return nil
}

//Log 打印日志
func Log(level LOG_LEVEL, v ...interface{}) {

}

//Write 写入日志
func (o *Handler) Write(level string, ts string, s string, v ...interface{}) {
	if LogOutput == OUTPUT_CSV {
		data := []string{level, ts, fmt.Sprintf("%v", v)}
		CsvWriter.Write(data)
		CsvWriter.Flush()
	} else {
		var a []interface{}
		a = append(a, ts)
		a = append(a, v...)
		LogWriter.SetPrefix(level)
		LogWriter.Printf("%-26s "+s, a...)
	}
}
