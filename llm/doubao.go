package llm

import (
	"context"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model/responses"
	"hello-ai-agent-go/config"
)

func DouBao(content string) (string, error) {

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
	l := len(resp.Output)
	return resp.Output[l-1].GetOutputMessage().Content[0].GetText().Text, err
}
