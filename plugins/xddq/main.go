package news

import (
	"path/filepath"
	"wxbot/engine/control"
	"wxbot/engine/robot"
)

func init() {
	engine := control.Register("xddq", &control.Options{
		Alias: "寻道",
		Help: "指令:\n" +
			"* 精怪轮回",
		DataFolder: "xddq",
	})
	engine.OnRegex(`^寻道(.+)轮回$`).SetBlock(true).Handle(func(ctx *robot.Ctx) {
		lunHui := ctx.State["regex_matched"].([]string)[1]
		var imgCache string
		if "精怪" == lunHui {
			imgCache = filepath.Join(engine.GetCacheFolder(), "jg.jpg")
		} else if "灵兽" == lunHui {
			imgCache = filepath.Join(engine.GetCacheFolder(), "ls.jpg")
		} else if "神通" == lunHui {
			imgCache = filepath.Join(engine.GetCacheFolder(), "st.jpg")
		} else {
			return
		}
		ctx.ReplyImage("local://" + imgCache)
	})
}
