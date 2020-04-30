//Package console 控制台模块.
//	可以在控制台输入指令来控制程序
package console

import (
	"bufio"
	"kada/server/core"
	"kada/server/log"
	"os"
	"strings"
)

var (
	_callbacks map[string]CallBack
)

type CallBack func(...string)

func init() {
	_callbacks = make(map[string]CallBack)
}

// 注册控制台消息
func Register(cmd string, callback CallBack) {
	_callbacks[cmd] = callback
}

// 监听控制台消息
func Listen() error {
	defer core.Panic()
	log.Signal("[console] wait listening cmd ...")

	reader := bufio.NewReader(os.Stdin)
	for {
		cmd, _, err := reader.ReadLine()
		if err != nil {
			return err
		}

		cmds := strings.Split(string(cmd), " ")
		log.Info("[console] cmd: %v", cmds)

		if cmds[0] == "over" {
			return nil
		}

		fun, ok := _callbacks[cmds[0]]
		if !ok {
			log.Warn("[console] no cmd: %s", cmds[0])
			continue
		}
		go fun(cmds[1:]...)
	}
}
