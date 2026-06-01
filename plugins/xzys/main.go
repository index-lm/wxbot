package xzys

import (
	"fmt"
	"github.com/imroc/req/v3"
	"wxbot/engine/control"
	"wxbot/engine/pkg/log"
	"wxbot/engine/robot"
)

func init() {
	engine := control.Register("xzys", &control.Options{
		Alias: "星座运势",
		Help: "描述:\n" +
			"星座运势，让我们一起看看吧\n\n" +
			"指令:\n" +
			"* 今日[] -> 获取肯德基疯狂星期四骚话",
	})

	engine.OnSuffix("运势").SetBlock(true).Handle(handle)
}

func handle(ctx *robot.Ctx) {
	xz := fmt.Sprintf("%s", ctx.State["args"])
	log.Printf("%s", xz)
	var data apiResponse
	xzCode := convertXZ(xz)
	log.Printf("%s", xzCode)
	if xzCode == "" {
		return
	}
	api := fmt.Sprintf("https://api.vvhan.com/api/horoscope?type=%s&time=today", xzCode)
	if err := req.C().Get(api).Do().Into(&data); err != nil {
		log.Errorf("获取失败: %v", err)
		return
	}
	if !data.Success {
		return
	}
	rs := buildRs(data)
	err := ctx.ReplyText(rs)
	// 将结果放到匹配队列，触发其它插件
	if err == nil {
		return
	}
}

func buildRs(data apiResponse) string {
	console := "今日运势\n"
	console += fmt.Sprintf("宜:  %s\n", data.Data.Todo.Yi)
	console += fmt.Sprintf("忌:  %s\n", data.Data.Todo.Ji)
	console += fmt.Sprintf("幸运指数:  %s\n", data.Data.Luckynumber)
	console += fmt.Sprintf("综合:\n  %s\n", data.Data.Fortunetext.All)
	console += fmt.Sprintf("爱情:\n  %s\n", data.Data.Fortunetext.Love)
	console += fmt.Sprintf("工作:\n  %s\n", data.Data.Fortunetext.Work)
	console += fmt.Sprintf("财运:\n  %s\n", data.Data.Fortunetext.Money)
	console += fmt.Sprintf("健康:\n  %s\n", data.Data.Fortunetext.Health)
	return console
}

func convertXZ(xz string) string {
	if xz == "白羊" {
		return "aries"
	} else if xz == "金牛" {
		return "taurus"
	} else if xz == "双子" {
		return "gemini"
	} else if xz == "巨蟹" {
		return "cancer"
	} else if xz == "狮子" {
		return "leo"
	} else if xz == "处女" {
		return "virgo"
	} else if xz == "天秤" {
		return "libra"
	} else if xz == "天蝎" {
		return "scorpio"
	} else if xz == "射手" {
		return "sagittarius"
	} else if xz == "摩羯" {
		return "capricorn"
	} else if xz == "水瓶" {
		return "aquarius"
	} else if xz == "双鱼" {
		return "pisces"
	} else {
		return ""
	}
}
