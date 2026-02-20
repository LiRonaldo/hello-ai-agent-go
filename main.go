package main

import (
	"fmt"
	_ "hello-ai-agent-go/config"
	"hello-ai-agent-go/llm"
)

func main() {
	content, err := llm.DouBao("介绍下高密")
	if err != nil {
		panic("我需要休息，不会思考了")
	}
	fmt.Println(content)
}
