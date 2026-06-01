package tbbi

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
	engine := control.Register("tbbi", &control.Options{
		Alias: "淘宝买家秀",
		Help: "指令:\n" +
			"* 淘宝买家秀 -> 淘宝买家秀",
		DataFolder: "tbbi",
	})
	engine.OnFullMatch("淘宝买家秀").SetBlock(true).Handle(func(ctx *robot.Ctx) {
		var resp ApiResponse
		u := rand.Uint64()
		var api = "https://api.vvhan.com/api/tao?type=json"
		if err := req.C().Get(api).Do().Into(&resp); err != nil {
			return
		}
		if err := ctx.ReplyImage(resp.Pic); err != nil {
			log.Errorf("[plmm] 发送图片失败: %v", err)
		}
		is := fmt.Sprintf("%d", u)
		imgCache := filepath.Join(engine.GetCacheFolder(), time.Now().Local().Format("20060102150405")+is+".jpg")
		// 下载图片
		if err := req.C().Get(resp.Pic).SetOutputFile(imgCache).Do().Err; err != nil {
			log.Errorf("[tbbi]下载图片失败: %v", err)
		}
	})
}
