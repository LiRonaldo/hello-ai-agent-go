package llm

import (
	"context"
	"fmt"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model/responses"
	"hello-ai-agent-go/config"
)

func DouBao(content string) (string, error) {
	fmt.Println("正在调用豆包大模型。。。。。")
	client := arkruntime.NewClientWithApiKey(
		// Get API Key：https://console.volcengine.com/ark/region:ark+cn-beijing/apikey
		config.Cfg.DouBao.ApiKey,
		arkruntime.WithBaseUrl(config.Cfg.DouBao.BaseUrl),
	)
	ctx := context.Background()

	resp, err := client.CreateResponses(ctx, &responses.ResponsesRequest{
		Model: config.Cfg.DouBao.ModelID,
		Input: &responses.ResponsesInput{Union: &responses.ResponsesInput_StringValue{StringValue: content}}, // Replace with your prompt
		// Thinking: &responses.ResponsesThinking{Type: responses.ThinkingType_disabled.Enum()}, // Manually disable deep thinking
	})
	if err != nil {
		return "", err
	}
	if resp.Status.String() == "completed" {
		fmt.Println("思考完成")
		fmt.Println("")
		fmt.Println("")
	}
	l := len(resp.Output)
	return resp.Output[l-1].GetOutputMessage().Content[0].GetText().Text, err
}
