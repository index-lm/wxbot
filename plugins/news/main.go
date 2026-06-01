package news

import (
	"fmt"
	"github.com/imroc/req/v3"
	"wxbot/engine/control"
	"wxbot/engine/robot"
)

func init() {
	engine := control.Register("news", &control.Options{
		Alias: "百度百科",
		Help: "指令:\n" +
			"* 百度百科 [查询内容]",
	})

	engine.OnFullMatch(`黑丝`).SetBlock(true).Handle(func(ctx *robot.Ctx) {
		if data, err := getBaiKe("黑丝", 0); err == nil {
			if data == nil {
				ctx.ReplyText("没查到该百科含义")
			} else {
				ctx.ReplyText("🏷黑丝" + ":\n" + fmt.Sprintf("%s\n🔎 摘要: %s\n©️ 版权: %s", data.Desc, data.Abstract, data.Copyrights))
			}
		} else {
			ctx.ReplyText("查询失败，这一定不是bug🤔")
		}
	})
}

type apiResponse struct {
	Key        string `json:"key"`
	Desc       string `json:"desc"`
	Abstract   string `json:"abstract"`
	Copyrights string `json:"copyrights"`
}

func getBaiKe(keyword string, i int) (*apiResponse, error) {
	if i >= 5 {
		return nil, nil
	}
	i++
	var data apiResponse
	api := "https://baike.baidu.com/api/openapi/BaikeLemmaCardApi?appid=379020&bk_length=1000&bk_key=" + keyword
	if err := req.C().Get(api).Do().Into(&data); err != nil {
		getBaiKe(keyword, i)
	}
	if len(data.Abstract) == 0 {
		getBaiKe(keyword, i)
	}
	return &data, nil
}
