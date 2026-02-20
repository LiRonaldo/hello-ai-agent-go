package tools

import (
	"encoding/json"
	"fmt"
	"hello-ai-agent-go/config"
	"hello-ai-agent-go/utils"
)

type messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type baiDuApiResp struct {
	Code       string      `json:"code"`
	RequestId  string      `json:"request_id"`
	Message    string      `json:"message"`
	References []reference `json:"references"`
}

type reference struct {
	Content   string      `json:"content"`    // 内容描述
	Date      string      `json:"date"`       // 发布时间
	Icon      interface{} `json:"icon"`       // 图标（null 用 interface{} 兼容）
	ID        int         `json:"id"`         // 唯一标识
	Image     interface{} `json:"image"`      // 图片（null 用 interface{} 兼容）
	Title     string      `json:"title"`      // 标题
	Type      string      `json:"type"`       // 类型（如 web）
	URL       string      `json:"url"`        // 链接地址
	Video     interface{} `json:"video"`      // 视频（null 用 interface{} 兼容）
	WebAnchor string      `json:"web_anchor"` // 网页锚点/副标题
}

func Search(query string) (string, error) {
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header["X-Appbuilder-Authorization"] = "Bearer " + config.Cfg.BaiDu.ApiKey
	param := []messages{
		{
			Role:    "user",
			Content: query,
		},
	}
	buf, err := json.Marshal(param)
	if err != nil {

		return "", err
	}
	resp, err := utils.Post(config.Cfg.BaiDu.BaseUrl, buf, map[string]string{})
	if err != nil {
		return "", err
	}
	fmt.Println(resp)
	defer resp.Body.Close()
	var r baiDuApiResp
	err = json.NewDecoder(resp.Body).Decode(&r)
	return "", nil
}
