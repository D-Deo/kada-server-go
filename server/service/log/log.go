//Package log 日志模块.
//	向控制台和文件输出日志信息
//	[DEBUG]调试类型日志只向控制台打印
//	[SIGN]信号类型的日志，打印级别是最高的
package log

import (
	"encoding/csv"
	"fmt"
	"kada/server/core"
	"kada/server/utils/config"
	"log"
	"os"
	"strings"
)

type LevelType int

const (
	_ LevelType = iota
	LvSignal
	LvError
	LvWarn
	LvInfo
	LvDebug
	LvCrash
)

const (
	OutputLog = "log"
	OutputCsv = "csv"
)

var (
	//LogLevel 控制台日志显示等级
	Level LevelType
	//LogOutput 输出类型
	Output string
	
	//Logger 日志控制器
	Logger *log.Logger
	
	//Writer 文件控制器
	Writer *csv.Writer
)

func Load(filename string) error {
	if filename == "" {
		filename = "kada"
	}
	
	Level = LvDebug
	if level, err := config.ToInt(config.Get(config.Logger, config.LoggerLevel)); err == nil {
		Level = LevelType(level)
	}
	
	Output = config.GetWithDef(config.Logger, config.LoggerOutput, OutputLog)
	filename += "." + Output
	
	var f *os.File
	var err error
	
	newFile := false
	
	if core.CheckFileIsExist(filename) { //如果文件存在
		f, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666) //打开文件
	} else {
		f, err = os.Create(filename)  //创建文件
		f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
		newFile = true
	}
	
	if err != nil {
		return err
	}
	
	if Output == OutputCsv {
		Writer = csv.NewWriter(f)
		if newFile {
			if head, err := config.Get(config.Logger, config.LoggerCsvHead); err != nil {
				Writer.Write(strings.Split(head, ","))
				Writer.Flush()
			}
			newFile = false
		}
	} else {
		// Logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
		Logger = log.New(f, "", 0)
	}
	
	return nil
}

//Write 写入日志
func Write(level string, ts string, s string, v ...interface{}) {
	if Output == OutputCsv {
		data := []string{level, ts, fmt.Sprintf("%v", v)}
		Writer.Write(data)
		Writer.Flush()
	} else {
		var a []interface{}
		a = append(a, ts)
		a = append(a, s)
		a = append(a, v...)
		Logger.SetPrefix(level)
		Logger.Printf(" %-26s %s", a...)
	}
}
