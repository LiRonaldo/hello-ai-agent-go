package agents

import (
	"fmt"
	"hello-ai-agent-go/constants"
	"hello-ai-agent-go/llm"
	"hello-ai-agent-go/utils"
)

type PlanAndSolveAgent struct {
	Question string
}

func NewPlanAndSolveAgent(question string) *PlanAndSolveAgent {
	return &PlanAndSolveAgent{
		Question: question,
	}
}

func (p *PlanAndSolveAgent) Run() {
	fmt.Println("我是一个PlanAndSolve类型的agent，我先将问题拆解成多步骤，在一步步计算")
	fmt.Println("用户的问题是:", p.Question)
	planList, err := plan(p.Question)
	if err != nil {
		fmt.Println("出问题了。。。。")
		return
	}
	executor(planList, p.Question)
}

// Plan 计划步骤
func plan(question string) ([]string, error) {
	prompt := fmt.Sprintf(constants.PlannerPromptTemplate, question)
	out, err := think(prompt)
	if err != nil {
		return nil, err
	}
	// 解析llm返回
	planList := parse(out)

	return planList, nil
}

// 执行
func executor(planList []string, question string) {
	fmt.Println("开始拆分步骤。。。。")
	history := ""
	outResp := ""
	for i, v := range planList {
		fmt.Println(fmt.Sprintf("当前正在执行第%d/%d步:%s", i+1, len(planList), v))
		out, err := think(fmt.Sprintf(constants.ExecutorPromptTemplate, question, planList, history, i+1))
		if err != nil {
			fmt.Println("出问题了。。。。。")
			return
		}
		history += fmt.Sprintf("第%d步: %s\n", i+1, out)
		outResp = out
		fmt.Println(fmt.Sprintf("已执行完第%d步,结果:%s", i+1, outResp))
	}
	fmt.Println(fmt.Sprintf("最终结果是:%s", outResp))
}

func think(prompt string) (string, error) {
	return llm.DouBao(prompt)
}

// ```go
// ["根据周一卖出15个苹果的已知条件，计算周二的苹果销量（周二销量 = 15 × 2）", "根据步骤1得到的周二苹果销量，计算周三的苹果销量（周三销量 = 周二销量 - 5）", "将周一的15个苹果、步骤1的周二销量、步骤2的周三销量相加，得出三天总销量"]
// ```
func parse(out string) []string {
	regex := utils.NewRegexp(constants.PlannerMatch)
	buf := regex.FindAllStringSubmatch(out, -1)
	planList := make([]string, 0, len(buf))
	for _, v := range buf {
		planList = append(planList, v[1])
	}
	return planList
}
