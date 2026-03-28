package core

import (
	"fmt"
)

// 记忆
type Memory struct {
	Execution  []string // 执行的记忆
	Reflection []string // 反思的记忆
}

func NewMemory(maxIterations uint8) *Memory {
	return &Memory{
		Execution:  make([]string, 0, maxIterations),
		Reflection: make([]string, 0, maxIterations),
	}
}

func (m *Memory) PushExecution(content string) {
	m.Execution = append(m.Execution, content)
}
func (m *Memory) PushReflection(content string) {
	m.Reflection = append(m.Reflection, content)
}

func (m *Memory) PopExecution() string {
	return fmt.Sprintf("--- 上一轮尝试 (代码) ---\n%s", m.Execution[len(m.Execution)-1])
}
func (m *Memory) PopReflection() string {
	return fmt.Sprintf("--- 评审员反馈 ---\n%s", m.Reflection[len(m.Reflection)-1])
}
func (m *Memory) GetLastExecution() string {
	return m.Execution[len(m.Execution)-1]
}
