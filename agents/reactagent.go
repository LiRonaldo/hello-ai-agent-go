package agents

import (
	"fmt"
	"hello-ai-agent-go/constants"
	"hello-ai-agent-go/llm"
	"hello-ai-agent-go/tools"
	"hello-ai-agent-go/utils"
	"strings"
)

/*
  - ReAct Agent
    思考与行动是相辅相成的

Thought (思考)： 这是智能体的“内心独白”。它会分析当前情况、分解任务、制定下一步计划，或者反思上一步的结果。
Action (行动)： 这是智能体决定采取的具体动作，通常是调用一个外部工具，例如 Search['华为最新款手机']。
Observation (观察)： 这是执行Action后从外部工具返回的结果，例如搜索结果的摘要或API的返回值。

智能体将不断重复这个 Thought -> Action -> Observation 的循环，将新的观察结果追加到历史记录中，
形成一个不断增长的上下文，直到它在Thought中认为已经找到了最终答案，
然后输出结果。这个过程形成了一个强大的协同效应：推理使得行动更具目的性，而行动则为推理提供了事实依据。
*/
type ReactAgent struct {
	history     string // 记录历史
	currentStep int    // 当前的思考步数
	maxStep     int    // 最多思考的步数
	question    string
}

func NewReActAgent(question string) *ReactAgent {
	return &ReactAgent{
		currentStep: 0,
		maxStep:     5,
		question:    question,
	}
}

func (reAct *ReactAgent) Run() {
	fmt.Println("管大海，您好，我是大海，能为您服务是我的荣幸!")
	fmt.Printf("事事真多，问我:%s\n", reAct.question)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("你别叫管大海了,你叫管扒皮吧。我累了,需要休息。")
		}
	}()
	for reAct.currentStep < reAct.maxStep {
		reAct.currentStep++
		msg := fmt.Sprintf(constants.ReactPromptTemplate, constants.ToolsSearch, constants.ToolsSearch, reAct.question, reAct.question, reAct.history)
		content, err := llm.DouBao(msg)
		if err != nil {
			fmt.Println("管大海，我刚才走神了！")
			break
		}
		thoughtRex := utils.NewRegexp(constants.ThoughtMatch)
		thoughtList := thoughtRex.FindStringSubmatch(content)
		// thought
		_ = thoughtList[1]
		fmt.Println("管扒皮,别催了,我正在思考")
		actionRex := utils.NewRegexp(constants.ActionMatch)
		actionList := actionRex.FindStringSubmatch(content)
		//action
		action := actionList[1]
		// 如果找到了最终的答案
		if strings.HasPrefix(action, constants.DouBaoFinish) {
			finishRex := utils.NewRegexp(constants.FinishMatch)
			answer := finishRex.FindStringSubmatch(action)
			fmt.Println("管大海,我找到了答案:")
			fmt.Println("-------------")
			fmt.Print(wrapByPeriod(answer[1]))
			fmt.Println("-------------")
			fmt.Println("您好满意吗?快夸夸我！")

			return
		}
		// 没有找到最终答案,需要调用search工具
		toolRex := utils.NewRegexp(constants.ToolMatch)
		toolList := toolRex.FindStringSubmatch(action)
		searchInput := toolList[2]
		searchOutput, err := tools.Search(searchInput)
		if err != nil {
			fmt.Println("我调用搜索工具失败。。。。")
			continue
		}
		reAct.history = fmt.Sprintf("Action:%s%s", action, searchOutput)
	}
}

// 按中文句号自动换行（如需处理英文句号，替换为 "." 即可）
func wrapByPeriod(text string) string {
	// 替换所有中文句号为“句号+换行符”
	// 若要处理英文句号：strings.ReplaceAll(text, ".", ".\n")
	return strings.ReplaceAll(text, "。", "。\n")
}
