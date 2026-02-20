package tools

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"hello-ai-agent-go/utils"
	"testing"
)

func Test_Search(t *testing.T) {
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header["X-Appbuilder-Authorization"] = "Bearer " + "bce-v3/ALTAK-hOfKGbMFkzBMeVh4nSbBy/1f0b33ef3b0437d8128a5779c6995471e4b32498"
	msg := []messages{
		{
			Role:    "user",
			Content: "北京的天气",
		},
	}
	params := make(map[string]interface{})
	params["messages"] = msg
	buf, err := json.Marshal(params)
	if err != nil {
		assert.NoError(t, err)
	}
	res, err := utils.Post("https://qianfan.baidubce.com/v2/ai_search/web_search", buf, header)
	assert.NoError(t, err)
	defer res.Body.Close()
	var r baiDuApiResp
	err = json.NewDecoder(res.Body).Decode(&r)
	assert.NoError(t, err)
	fmt.Println(res)
}
