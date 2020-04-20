package log

import (
	"fmt"
	"github.com/issue9/term/colors"
	"time"
)

type Handler int

type Args struct {
	Level  LevelType // 日志显示等级
	Output string    // 输出类型
}

func (o *Handler) Log(args *Args, back *int) error {
	ts := time.Now().Format("2006-01-02 15:04:05.999999") //设定时间格式
	
	switch args.Level {
	case LvDebug:
		if Level >= args.Level {
			if _, err := colors.Println(colors.Cyan, colors.Default, fmt.Sprintf("%s %-26s %s", "[D]", ts, args.Output)); err != nil {
				return err
			}
			Write("[D]", ts, args.Output)
		}
	case LvInfo:
		if Level >= args.Level {
			if _, err := colors.Println(colors.White, colors.Default, fmt.Sprintf("%s %-26s %s", "[I]", ts, args.Output)); err != nil {
				return err
			}
		}
		Write("[I]", ts, args.Output)
	case LvSignal:
		if Level >= args.Level {
			if _, err := colors.Println(colors.Green, colors.Default, fmt.Sprintf("%s %-26s %s", "[S]", ts, args.Output)); err != nil {
				return err
			}
		}
		Write("[S]", ts, args.Output)
	case LvWarn:
		if Level >= args.Level {
			if _, err := colors.Println(colors.Yellow, colors.Default, fmt.Sprintf("%s %-26s %s", "[W]", ts, args.Output)); err != nil {
				return err
			}
		}
		Write("[W]", ts, args.Output)
	case LvError:
		if Level >= args.Level {
			if _, err := colors.Println(colors.Red, colors.Default, fmt.Sprintf("%s %-26s %s", "[E]", ts, args.Output)); err != nil {
				return err
			}
		}
		Write("[E]", ts, args.Output)
	case LvCrash:
		Write("[P]", ts, args.Output)
	default:
		break
	}
	return nil
}
