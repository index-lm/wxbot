package acgimg

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
	engine := control.Register("acgimg", &control.Options{
		Alias: "二次元图片",
		Help: "指令:\n" +
			"* 二次元 -> 二次元图片",
		DataFolder: "acgimg",
	})
	engine.OnFullMatch("二次元").SetBlock(true).Handle(func(ctx *robot.Ctx) {
		var resp ApiResponse
		u := rand.Uint64()
		var api = "https://api.vvhan.com/api/acgimg?type=json"
		if err := req.C().Get(api).Do().Into(&resp); err != nil {
			return
		}
		if err := ctx.ReplyImage(resp.Imgurl); err != nil {
			log.Errorf("[acgimg] 发送图片失败: %v", err)
		}
		is := fmt.Sprintf("%d", u)
		imgCache := filepath.Join(engine.GetCacheFolder(), time.Now().Local().Format("20060102150405")+is+".jpg")
		// 下载图片
		if err := req.C().Get(resp.Imgurl).SetOutputFile(imgCache).Do().Err; err != nil {
			log.Errorf("[acgimg]下载图片失败: %v", err)
		}
	})
	engine.OnFullMatch("LOL").SetBlock(true).Handle(func(ctx *robot.Ctx) {
		var resp ApiResponse
		u := rand.Uint64()

		if err := ctx.ReplyImage("https://api.vvhan.com/api/lolskin"); err != nil {
			log.Errorf("[acgimg] 发送图片失败: %v", err)
		}
		is := fmt.Sprintf("%d", u)
		imgCache := filepath.Join(engine.GetCacheFolder()+"/lol", time.Now().Local().Format("20060102150405")+is+".jpg")
		// 下载图片
		if err := req.C().Get(resp.Imgurl).SetOutputFile(imgCache).Do().Err; err != nil {
			log.Errorf("[acgimg]下载图片失败: %v", err)
		}
	})
	engine.OnFullMatch("随机风景").SetBlock(true).Handle(func(ctx *robot.Ctx) {
		var resp ApiResponse
		u := rand.Uint64()

		if err := ctx.ReplyImage("https://api.vvhan.com/api/view"); err != nil {
			log.Errorf("[acgimg] 发送图片失败: %v", err)
		}
		is := fmt.Sprintf("%d", u)
		imgCache := filepath.Join(engine.GetCacheFolder()+"/fj", time.Now().Local().Format("20060102150405")+is+".jpg")
		// 下载图片
		if err := req.C().Get(resp.Imgurl).SetOutputFile(imgCache).Do().Err; err != nil {
			log.Errorf("[acgimg]下载图片失败: %v", err)
		}
	})
}
