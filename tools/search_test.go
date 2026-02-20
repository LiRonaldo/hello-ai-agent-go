package tools

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"hello-ai-agent-go/config"
	"hello-ai-agent-go/utils"
	"testing"
)

func Test_Search(t *testing.T) {
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header["X-Appbuilder-Authorization"] = "Bearer " + config.Cfg.BaiDu.ApiKey
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
	res, err := utils.Post(config.Cfg.BaiDu.BaseUrl, buf, header)
	assert.NoError(t, err)
	defer res.Body.Close()
	var r baiDuApiResp
	err = json.NewDecoder(res.Body).Decode(&r)
	assert.NoError(t, err)
	fmt.Println(res)
}
