package sjyztx

import (
	"fmt"
	"github.com/imroc/req/v3"
	"math/rand"
	"path/filepath"
	"time"
	"wxbot/engine/control"
	"wxbot/engine/pkg/log"
	"wxbot/engine/robot"
)

func init() {
	engine := control.Register("sjyztx", &control.Options{
		Alias: "随机一张头像",
		Help: "描述:\n" +
			"随机一张头像，让我们一起看看吧\n\n" +
			"指令:\n" +
			"* 随机头像 -> 随机一张头像",
		DataFolder: "sjyztx",
	})

	engine.OnFullMatch("随机头像").SetBlock(true).Handle(func(ctx *robot.Ctx) {
		var resp ApiResponse
		api := "https://api.vvhan.com/api/avatar?type=json"
		if err := req.C().SetBaseURL(api).Get().Do().Into(&resp); err != nil {
			return
		}
		if !resp.Success {
			return
		}
		if err := ctx.ReplyImage(resp.Avatar); err != nil {
			log.Errorf("[sjyztx] 发送图片失败: %v", err)
		}
		is := fmt.Sprintf("%d", rand.Uint64())
		imgCache := filepath.Join(engine.GetCacheFolder(), time.Now().Local().Format("20060102150405")+is+".png")
		// 下载图片
		if err := req.C().Get(resp.Avatar).SetOutputFile(imgCache).Do().Err; err != nil {
			log.Errorf("[sjyztx]下载图片失败: %v", err)
		}
	})
}
