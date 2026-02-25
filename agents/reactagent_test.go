package agents

import (
	"fmt"
	"hello-ai-agent-go/constants"
	"hello-ai-agent-go/utils"
	"testing"
)

func Test_ReActAgent(t *testing.T) {
	agent := NewReActAgent("go语言最新版本")
	agent.Run()
}

func Test_Match(t *testing.T) {
	content := "Thought: 用户想了解华为最新款手机，由于手机产品更新迭代较快，需要通过搜索获取当前最新的信息。\nAction: Search[华为最新款手机]"
	thoughtRex := utils.NewRegexp(constants.ThoughtMatch)
	thoughtList := thoughtRex.FindStringSubmatch(content)
	thought := thoughtList[1]
	fmt.Println(thought)
	actionRex := utils.NewRegexp(constants.ActionMatch)
	actionList := actionRex.FindStringSubmatch(content)
	action := actionList[1]
	fmt.Println(action)
	toolRex := utils.NewRegexp(constants.ToolMatch)
	toolList := toolRex.FindStringSubmatch(action)
	fmt.Println(toolList[1])
	toolName := toolList[2]
	fmt.Println(toolName)
}
