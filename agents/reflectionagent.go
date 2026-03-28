package agents

import (
	"fmt"
	"hello-ai-agent-go/constants"
	"hello-ai-agent-go/core"
	"hello-ai-agent-go/llm"
	"strings"
)

/*
*
执行 -> 反思 -> 优化
*/

type ReflectionAgent struct {
	question      string
	maxIterations uint8
	memory        *core.Memory
}

// 默认3次反思
func NewReflectionAgent(question string, maxIterations uint8) *ReflectionAgent {
	return &ReflectionAgent{
		question:      question,
		maxIterations: maxIterations,
		memory:        core.NewMemory(maxIterations),
	}
}

func (agent *ReflectionAgent) Run() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("我累了,需要休息。")
		}
	}()
	fmt.Printf("---开始处理任务----\n%s\n", agent.question)
	msg := fmt.Sprintf(constants.InitialPromptTemplate, agent.question)
	code, err := llm.DouBao(msg)
	if err != nil {
		panic(err)
	}
	agent.memory.PushExecution(code)
	for i := uint8(0); i < agent.maxIterations; i++ {
		fmt.Printf("---第%d/%d次迭代----\n", i+1, agent.maxIterations)
		fmt.Printf("---开始反思---\n")
		leastCode := agent.memory.GetLastExecution()
		prompt := fmt.Sprintf(constants.ReflectPromptTemplate, agent.question, leastCode)
		feedback, err := llm.DouBao(prompt)
		if err != nil {
			panic(err)
		}
		agent.memory.PushReflection(feedback)
		if strings.Contains(feedback, constants.ReflectionAgentEnd) {
			fmt.Println("---结束反思,任务完成---")
			break
		}
		fmt.Println("---继续优化---")
		refinedCode, err := llm.DouBao(fmt.Sprintf(constants.RefinePromptTemplate, agent.question, leastCode, feedback))
		if err != nil {
			panic(err)
		}
		agent.memory.PushExecution(refinedCode)
	}
	fmt.Printf("---最终的代码---\n")
	fmt.Println(agent.memory.GetLastExecution())
}
