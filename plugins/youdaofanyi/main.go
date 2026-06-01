package youdaofanyi

import (
	"fmt"
	"net/url"

	"github.com/imroc/req/v3"

	"wxbot/engine/control"
	"wxbot/engine/robot"
)

func init() {
	engine := control.Register("youdaofanyi", &control.Options{
		Alias: "有道中英文互译",
		Help: "指令:\n" +
			"* 翻译 [内容]\n" +
			"* 有道翻译 [内容]\n",
	})

	engine.OnRegex(`(^有道翻译|^翻译) ?(.*?)$`).SetBlock(true).Handle(func(ctx *robot.Ctx) {
		word := ctx.State["regex_matched"].([]string)[2]
		if data, err := getFanYi(word); err == nil {
			if data == nil {
				ctx.ReplyText("我还不会，稍后尝试")
			} else {
				ctx.ReplyText(fmt.Sprintf("🔎 译文:\n %s", data.Result))
			}
		} else {
			ctx.ReplyText("查询失败，这一定不是bug🤔")
		}
	})
}

type apiResponse struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Name   string `json:"name"`
	Result string `json:"result"`
}

func getFanYi(keyword string) (*apiResponse, error) {
	var data apiResponse
	api := "https://api.qqsuu.cn/api/dm-ydfy?name=" + url.QueryEscape(keyword)
	if err := req.C().Get(api).Do().Into(&data); err != nil {
		return nil, err
	}
	if len(data.Result) == 0 {
		return nil, nil
	}
	return &data, nil
}
