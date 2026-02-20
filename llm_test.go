package main

import (
	"context"
	"fmt"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model/responses"
	"testing"
)

func Test_LLM(t *testing.T) {
	client := arkruntime.NewClientWithApiKey(
		// Get API Key：https://console.volcengine.com/ark/region:ark+cn-beijing/apikey
		"3502ebc7-9dc8-4062-8529-9650b4989d60",
		arkruntime.WithBaseUrl("https://ark.cn-beijing.volces.com/api/v3"),
	)
	ctx := context.Background()

	resp, err := client.CreateResponses(ctx, &responses.ResponsesRequest{
		Model: "doubao-seed-2-0-code-preview-260215",
		Input: &responses.ResponsesInput{Union: &responses.ResponsesInput_StringValue{StringValue: "你叫什么名字"}}, // Replace with your prompt
		// Thinking: &responses.ResponsesThinking{Type: responses.ThinkingType_disabled.Enum()}, // Manually disable deep thinking
	})
	if err != nil {
		fmt.Printf("response error: %v\n", err)
		return
	}
	fmt.Println(resp)
}
