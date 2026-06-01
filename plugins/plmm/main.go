package plmm

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
	engine := control.Register("plmm", &control.Options{
		Alias: "漂亮妹妹",
		Help: "指令:\n" +
			"* 漂亮妹妹 -> 获取漂亮妹妹",
		DataFolder: "plmm",
	})
	engine.OnFullMatch("漂亮妹妹").SetBlock(true).Handle(func(ctx *robot.Ctx) {
		var resp ApiResponse
		u := rand.Uint64()
		var api string
		if u%2 == 0 {
			api = "https://api.vvhan.com/api/wallpaper/mobileGirl?type=json"
		} else {
			api = "https://api.vvhan.com/api/wallpaper/pcGirl?type=json"
		}
		if err := req.C().SetBaseURL(api).Get().Do().Into(&resp); err != nil {
			return
		}
		if !resp.Success {
			return
		}
		if err := ctx.ReplyImage(resp.Url); err != nil {
			log.Errorf("[plmm] 发送图片失败: %v", err)
		}
		is := fmt.Sprintf("%d", u)
		imgCache := filepath.Join(engine.GetCacheFolder(), time.Now().Local().Format("20060102150405")+is+".png")
		// 下载图片
		if err := req.C().Get(resp.Url).SetOutputFile(imgCache).Do().Err; err != nil {
			log.Errorf("[zaoBao]下载图片失败: %v", err)
		}
	})
}
