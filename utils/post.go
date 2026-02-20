package utils

import (
	"net/http"
	"strings"
)

func Post(url string, args []byte, headers map[string]string) (*http.Response, error) {
	buf := strings.NewReader(string(args))
	request, err := http.NewRequest(http.MethodPost, url, buf)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		request.Header.Set(k, v)
	}
	// 创建一个 HTTP 客户端
	client := &http.Client{}
	// 发送请求
	return client.Do(request)

}
