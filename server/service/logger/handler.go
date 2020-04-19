package logger

import (
	"encoding/csv"
	"fmt"
	"github.com/issue9/term/colors"
	"kada/server/core"
	"kada/server/utils/config"
	"log"
	"os"
	"strings"
	"time"
)

type Handler int

type LoadArgs struct {
	FileName string // 输出类型
}

func (o *Handler) Load(args *LoadArgs, back *int) error {
	if args.FileName == "" {
		args.FileName = "kada"
	}
	
	LogLevel = LOG_DEBUG
	if level, err := config.ToInt(config.Get(config.Logger, config.LoggerLevel)); err == nil {
		LogLevel = LOG_LEVEL(level)
	}
	
	LogOutput = config.GetWithDef(config.Logger, config.LoggerOutput, OUTPUT_LOG)
	args.FileName += "." + LogOutput
	
	var f *os.File
	var err error
	
	newFile := false
	
	if core.CheckFileIsExist(args.FileName) { //如果文件存在
		f, err = os.OpenFile(args.FileName, os.O_WRONLY|os.O_APPEND, 0666) //打开文件
	} else {
		f, err = os.Create(args.FileName) //创建文件
		f.WriteString("\xEF\xBB\xBF")     // 写入UTF-8 BOM
		newFile = true
	}
	
	if err != nil {
		return err
	}
	
	if LogOutput == OUTPUT_CSV {
		CsvWriter = csv.NewWriter(f)
		if newFile {
			if head, err := config.Get(config.Logger, config.LoggerCsvHead); err != nil {
				CsvWriter.Write(strings.Split(head, ","))
				CsvWriter.Flush()
			}
			newFile = false
		}
	} else {
		// LogWriter = log.New(f, "", log.Ldate|log.Lmicroseconds)
		LogWriter = log.New(f, "", 0)
	}
	
	return nil
}

type LogArgs struct {
	Level  LOG_LEVEL // 日志显示等级
	Output string    // 输出类型
}

func (o *Handler) Log(args *LogArgs, back *int) error {
	ts := time.Now().Format("2006-01-02 15:04:05.999999") //设定时间格式
	
	//s := ""
	//for i := 0; i < len(v); i++ {
	//	if i != 0 {
	//		s += " "
	//	}
	//	s += "%#v"
	//}
	//var a []interface{}
	//a = append(a, ts)
	//a = append(a, v...)
	
	switch args.Level {
	case LOG_DEBUG:
		if LogLevel >= args.Level {
			if _, err := colors.Println(colors.Cyan, colors.Default, fmt.Sprintf("%s %-26s %s", "[D]", ts, args.Output)); err != nil {
				return err
			}
			//Write("[D]", ts, args.Output)
		}
	case LOG_INFO:
		if LogLevel >= args.Level {
			if _, err := colors.Println(colors.White, colors.Default, fmt.Sprintf("%s %-26s %s", "[I]", ts, args.Output)); err != nil {
				return err
			}
		}
		//Write("[I]", ts, args.Output)
	case LOG_SIGNAL:
		if LogLevel >= args.Level {
			if _, err := colors.Println(colors.Green, colors.Default, fmt.Sprintf("%s %-26s %s", "[S]", ts, args.Output)); err != nil {
				return err
			}
		}
		//Write("[S]", ts, args.Output)
	case LOG_WARN:
		if LogLevel >= args.Level {
			if _, err := colors.Println(colors.Yellow, colors.Default, fmt.Sprintf("%s %-26s %s", "[W]", ts, args.Output)); err != nil {
				return err
			}
		}
		//Write("[W]", ts, args.Output)
	case LOG_ERROR:
		if LogLevel >= args.Level {
			if _, err := colors.Println(colors.Red, colors.Default, fmt.Sprintf("%s %-26s %s", "[E]", ts, args.Output)); err != nil {
				return err
			}
		}
		//Write("[E]", ts, args.Output)
	case LOG_CRASH:
		//Write("[P]", ts, args.Output)
	default:
		break
	}
	return nil
}
