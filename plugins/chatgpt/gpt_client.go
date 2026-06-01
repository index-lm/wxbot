package chatgpt

import (
	"fmt"
	"github.com/imroc/req/v3"
	util_jwt "wxbot/engine/pkg/jwt"
	"wxbot/engine/pkg/log"
)

type GlmReq struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	//RequestId   string  `json:"request_id"`
	//DoSample    bool    `json:"do_sample"`
	Stream bool `json:"stream"`
	//Temperature float32 `json:"temperature"`
	//TopP        float32 `json:"top_p"`
	//MaxTokens   int     `json:"max_tokens"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GlmRes struct {
	Id      string    `json:"id"`
	Created int64     `json:"created"`
	Model   string    `json:"model"`
	Choices []Choices `json:"choices"`
	Usage   Usage     `json:"usage"`
}

type Choices struct {
	Index        int     `json:"index"`
	FinishReason string  `json:"finish_reason"`
	Message      Message `json:"message"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func sendArr(messages []Message) string {
	api := "https://open.bigmodel.cn/api/paas/v4/chat/completions"

	rpcReq := GlmReq{
		Model:    "glm-4",
		Stream:   false,
		Messages: messages,
	}
	rpcRes := new(GlmRes)
	token, err := util_jwt.Jwt.CreateToken()
	if err != nil {
		log.Errorf("创建jwt失败:%v", err)
	}
	fmt.Println("开始调用gpt")
	if err = req.C().Post(api).SetBodyJsonMarshal(&rpcReq).SetHeader("Authorization", "Bearer "+token).Do().Into(&rpcRes); err != nil {
		log.Errorf("获取失败: %v", err)
	}
	if rpcRes.Id != "" {
		return rpcRes.Choices[0].Message.Content
	} else {
		return "获取失败"
	}

}

func initMessage(msg string) *[]Message {
	return &[]Message{
		{
			Role:    "user",
			Content: "你现在是智能助手." + msg,
		},
	}
}
