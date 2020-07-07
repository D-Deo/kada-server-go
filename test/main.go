package main

import (
	"fmt"
	"github.com/fatih/color"
	"kada/server/console"
	"kada/server/core"
	"kada/server/log"
	"kada/server/utils/config"
	"kada/test/user"
	"runtime"
	"syscall"
)

var (
	kernel32    = syscall.NewLazyDLL(`kernel32.dll`)
	proc        = kernel32.NewProc(`SetConsoleTextAttribute`)
	CloseHandle = kernel32.NewProc(`CloseHandle`)

	// 给字体颜色对象赋值
	FontColor Color = Color{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
)

type Color struct {
	black        int // 黑色
	blue         int // 蓝色
	green        int // 绿色
	cyan         int // 青色
	red          int // 红色
	purple       int // 紫色
	yellow       int // 黄色
	light_gray   int // 淡灰色（系统默认值）
	gray         int // 灰色
	light_blue   int // 亮蓝色
	light_green  int // 亮绿色
	light_cyan   int // 亮青色
	light_red    int // 亮红色
	light_purple int // 亮紫色
	light_yellow int // 亮黄色
	white        int // 白色
}

// 输出有颜色的字体
func ColorPrint(s string, i int) {
	r1, r2, err := proc.Call(uintptr(syscall.Stdout), uintptr(i))
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	if _, _, err := CloseHandle.Call(r1, r2); err != nil {
		panic(err)
	}
}

func show(args ...string) {
	log.Warn("[main] server status NumCPU(%d) NumGoroutine(%d) NumCgoCall(%d)", runtime.NumCPU(), runtime.NumGoroutine(), runtime.NumCgoCall())
}

func test1(args ...string) {
	req := &user.CreateArgs{}
	req.UserName = "test"
	req.Password = "123123"
	req.IP = "127.0.0.1"
	req.Device = 1
	req.Phone = "12345678901"
	req.Token = "token_1"

	var ret user.CreateBack
	if err := user.Call("Create", req, &ret); err != nil {
		log.Panic("[test2] user create err: %v", err)
	}

	log.Info("[test2] user create ret: %v", ret.User)
}

func test2(args ...string) {
	color.Set(color.FgBlue)
	println("blue")
	color.Set(color.FgWhite)
	println("white")
	color.Set(color.FgRed)
	println("red")
	defer color.Unset()
	//ColorPrint("红色", FontColor.red)
	//ColorPrint("蓝色", FontColor.blue)
	//ColorPrint("白色", FontColor.white)
	//ColorPrint("灰色", FontColor.light_gray)
	//c := colors.New(colors.Normal, colors.Yellow, colors.Blue)
	//c.Println("123123")
	//colors.Println(colors.Normal, colors.Yellow, colors.Default, "TEST2")
}

func main() {
	defer core.Panic()

	// 利用cpu多核来控制go的协程
	ncpu := runtime.NumCPU()
	ngoc := runtime.GOMAXPROCS(ncpu)

	// 初始化配置
	if err := config.Load("test"); err != nil {
		log.Panic(err)
	}

	log.Signal("[main] monitor startup CPU: %d GOMAXPROC: %d", ncpu, ngoc)

	console.Register("show", show)
	console.Register("test1", test1)
	console.Register("test2", test2)
	console.Listen()
}
