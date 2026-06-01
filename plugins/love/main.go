package sjyjyy

import (
	"github.com/imroc/req/v3"
	"wxbot/engine/control"
	"wxbot/engine/pkg/log"
	"wxbot/engine/robot"
)

func init() {
	engine := control.Register("love", &control.Options{
		Alias: "随机情话",
		Help: "描述:\n" +
			"随机情话，让我们一起看看吧\n\n" +
			"指令:\n" +
			"* 早安 -> 情话" +
			"* 午安 -> 情话" +
			"* 晚安 -> 情话",
	})

	engine.OnFullMatch("早安").SetBlock(true).Handle(handle1)
	engine.OnFullMatch("午安").SetBlock(true).Handle(handle2)
	engine.OnFullMatch("晚安").SetBlock(true).Handle(handle3)
	engine.OnFullMatch("随机情话").SetBlock(true).Handle(handle4)
}

func handle1(ctx *robot.Ctx) {
	var rpcRes apiResponse
	api := "https://api.vvhan.com/api/text/love?type=json"
	if err := req.C().Get(api).Do().Into(&rpcRes); err != nil {
		log.Errorf("获取失败: %v", err)
		return
	}
	res := rpcRes.Data.Content + "\n[咖啡]早上好~"
	err := ctx.ReplyText(res)
	// 将结果放到匹配队列，触发其它插件
	if err == nil {
		return
	}
}
func handle2(ctx *robot.Ctx) {
	var rpcRes apiResponse
	api := "https://api.vvhan.com/api/text/love?type=json"
	if err := req.C().Get(api).Do().Into(&rpcRes); err != nil {
		log.Errorf("获取失败: %v", err)
		return
	}
	res := rpcRes.Data.Content + "\n[太阳]中午好~"
	err := ctx.ReplyText(res)
	// 将结果放到匹配队列，触发其它插件
	if err == nil {
		return
	}
}
func handle3(ctx *robot.Ctx) {
	var rpcRes apiResponse
	api := "https://api.vvhan.com/api/text/love?type=json"
	if err := req.C().Get(api).Do().Into(&rpcRes); err != nil {
		log.Errorf("获取失败: %v", err)
		return
	}
	res := rpcRes.Data.Content + "\n[月亮]晚安~"
	err := ctx.ReplyText(res)
	// 将结果放到匹配队列，触发其它插件
	if err == nil {
		return
	}
}
func handle4(ctx *robot.Ctx) {
	var rpcRes apiResponse
	api := "https://api.vvhan.com/api/text/love?type=json"
	if err := req.C().Get(api).Do().Into(&rpcRes); err != nil {
		log.Errorf("获取失败: %v", err)
		return
	}
	res := rpcRes.Data.Content + ""
	err := ctx.ReplyText(res)
	// 将结果放到匹配队列，触发其它插件
	if err == nil {
		return
	}
}
