package sjyjyy

import (
	"github.com/imroc/req/v3"
	"wxbot/engine/control"
	"wxbot/engine/pkg/log"
	"wxbot/engine/robot"
)

func init() {
	engine := control.Register("sjyjyy", &control.Options{
		Alias: "随机一句一言",
		Help: "描述:\n" +
			"随机一句一言，让我们一起看看吧\n\n" +
			"指令:\n" +
			"* 随机一言 -> 获取肯德基疯狂星期四骚话",
	})

	engine.OnFullMatch("随机一言").SetBlock(true).Handle(handle)
	engine.OnFullMatch("随机骚话").SetBlock(true).Handle(handle2)
	engine.OnFullMatch("随机笑话").SetBlock(true).Handle(handle3)
}

func handle(ctx *robot.Ctx) {
	var data apiResponse
	api := "https://api.vvhan.com/api/ian?type=json"
	if err := req.C().Get(api).Do().Into(&data); err != nil {
		log.Errorf("获取失败: %v", err)
		return
	}
	err := ctx.ReplyText(data.Data.Vhan)
	// 将结果放到匹配队列，触发其它插件
	if err == nil {
		return
	}
}

func handle2(ctx *robot.Ctx) {
	var data apiResponse2
	api := "https://api.vvhan.com/api/sao?type=json"
	if err := req.C().Get(api).Do().Into(&data); err != nil {
		log.Errorf("获取失败: %v", err)
		return
	}
	err := ctx.ReplyText(data.Ishan)
	// 将结果放到匹配队列，触发其它插件
	if err == nil {
		return
	}
}
func handle3(ctx *robot.Ctx) {
	var data apiResponse3
	api := "https://api.vvhan.com/api/joke?type=json"
	if err := req.C().Get(api).Do().Into(&data); err != nil {
		log.Errorf("获取失败: %v", err)
		return
	}
	err := ctx.ReplyText(data.Joke)
	// 将结果放到匹配队列，触发其它插件
	if err == nil {
		return
	}
}
