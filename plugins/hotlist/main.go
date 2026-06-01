package news

import (
	"wxbot/engine/control"
	"wxbot/engine/robot"
)

func init() {
	engine := control.Register("news", &control.Options{
		Alias: "热榜",
		Help: "指令:\n" +
			"* []热榜",
	})
	engine.OnSuffix("热榜").SetBlock(true).Handle(handle)
}

func handle(ctx *robot.Ctx) {
	err := ctx.ReplyText("这是一个测试信息")
	// 将结果放到匹配队列，触发其它插件
	if err == nil {
		return
	}

}
