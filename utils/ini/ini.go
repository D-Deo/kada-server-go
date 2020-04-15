package ini

import (
	"errors"
	"fmt"
	"github.com/larspensjo/config"
	"strconv"
	"strings"
)

var (
	DefaultIni *Ini
	ErrNoValue = errors.New("config error, no value") //没有对应的值
)

// NewIni 创建新的配置对象
func NewIni() *Ini {
	return &Ini{}
}

// Load 加载配置文件
func Load(filename string) error {
	DefaultIni = NewIni()
	if err := DefaultIni.Load(filename); err != nil {
		return err
	}
	return nil
}

// 获取内容
func Get(title string, key string) (string, error) {
	return DefaultIni.Get(title, key)
}

// 获取内容，如果不存在则返回默认值
func GetWithDef(title string, key string, def string) string {
	return DefaultIni.GetWithDef(title, key, def)
}

// 转 int 输出
func ToInt(str string, err error) (int, error) {
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(str)
}

// Ini 配置对象
type Ini struct {
	cfg *config.Config
}

func (o *Ini) Load(filename string) error {
	if !strings.Contains(filename, ".ini") {
		filename += ".ini"
	}
	cfg, err := config.ReadDefault(filename)
	if err != nil {
		return fmt.Errorf("[ini] load (%s) err: %w", filename, err)
	}
	o.cfg = cfg
	return nil
}

// 获取内容
func (o *Ini) Get(title string, key string) (string, error) {
	if !o.cfg.HasOption(title, key) {
		return "", ErrNoValue
	}

	value, err := o.cfg.String(title, key)
	if err != nil {
		return "", err
	}

	return value, nil
}

// 获取内容，如果不存在则返回默认值
func (o *Ini) GetWithDef(title string, key string, def string) string {
	value, err := o.cfg.String(title, key)
	if err != nil {
		return def
	}
	return value
}
