package utils

import (
	"net/http"
	"strings"
)

func Get(url string, args map[string]string, headers map[string]string) (*http.Response, error) {
	sb := strings.Builder{}
	for k, v := range args {
		sb.WriteString(k)
		sb.WriteString("=")
		sb.WriteString(v)
		sb.WriteString("&")
	}
	if sb.Cap() > 0 {
		url = url + "?"
	}
	url += sb.String()
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		request.Header.Set(k, v)
	}
	// 创建一个 HTTP 客户端
	client := &http.Client{}
	// 发送请求
	return client.Do(request)
}
