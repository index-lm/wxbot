package localimagespider

import (
	"wxbot/engine/control"
	"wxbot/engine/robot"
	"wxbot/plugins/localimage"
)

func init() {
	engine := control.Register("localimagespider", &control.Options{
		Alias: "爬取图片到本地",
		Help: "指令:\n" +
			"* 抓取Cosplay作品\n" +
			"* 抓取Coser日常\n",
	})

	storageFolder := localimage.GetStorageFolder()

	engine.OnFullMatch("抓取Cosplay作品", robot.OnlyMe).SetBlock(true).Handle(func(ctx *robot.Ctx) {
		crawlCosplay(storageFolder)
	})
	engine.OnFullMatch("抓取Coser日常", robot.OnlyMe).SetBlock(true).Handle(func(ctx *robot.Ctx) {
		crawlCoser(storageFolder)
	})

}
