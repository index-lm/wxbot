package main

import (
	"time"

	"github.com/spf13/viper"
	"wxbot/engine/pkg/log"
	"wxbot/engine/pkg/net"
	"wxbot/engine/robot"
	"wxbot/framework/dean"
	"wxbot/framework/vlw"

	// 导入插件, 变更插件请查看README
	_ "wxbot/engine/plugins"
)

func main() {

	// E:\WeChat\[3.6.0.18]
	// E:\WeChat\wxhook\DaenWxHook.dll
	// callBackUrl=http://localhost:8867/wxbot/callback&port=8866&decryptlmg=1&imgData=C:\Users\78320\Documents\WeChat Files\
	v := viper.New()
	v.SetConfigFile("config.yaml")
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("[main] 读取配置文件失败: %s", err.Error())
	}
	c := robot.NewConfig()
	if err := v.Unmarshal(c); err != nil {
		log.Fatalf("[main] 解析配置文件失败: %s", err.Error())
	}

	f := robot.IFramework(nil)
	switch c.Framework.Name {
	case "Dean":
		f = robot.IFramework(dean.New(c.BotWxId, c.Framework.ApiUrl, c.Framework.ApiToken))
		if ipPort, err := net.CheckoutIpPort(c.Framework.ApiUrl); err == nil {
			if ping := net.PingConn(ipPort, time.Second*10); !ping {
				c.SetConnHookStatus(false)
				log.Warn("[main] 无法连接Dean框架，网络无法Ping通，请检查网络")
			}
		}
	case "VLW", "vlw":
		f = robot.IFramework(vlw.New(c.BotWxId, c.Framework.ApiUrl, c.Framework.ApiToken))
		if ipPort, err := net.CheckoutIpPort(c.Framework.ApiUrl); err == nil {
			if ping := net.PingConn(ipPort, time.Second*10); !ping {
				c.SetConnHookStatus(false)
				log.Warn("[main] 无法连接到VLW框架，网络无法Ping通，请检查网络")
			}
		}
	default:
		log.Fatalf("[main] 请在配置文件中指定机器人框架后再启动")
	}

	robot.Run(c, f)
}
