package setting

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
	"time"
)

// 调用app.ini配置的模块，数据库相关的配置读取在models中
// ini官方文档地址： https://ini.unknwon.io/docs/intro/getting_started

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read 'conf/app.ini': %v", err)
		os.Exit(1) //程序遇到了错误或异常情况,退出程序
	}
	LoadBase()
	LoadServer()
}

// 读取运行模式
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug") //MustString("debug")表示默认值
}

// 读取服务相关配置
func LoadServer() {
	sec, _ := Cfg.GetSection("server")
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8088)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second //转换成纳秒
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}
