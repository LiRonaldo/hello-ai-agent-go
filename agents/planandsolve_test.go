package agents

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlanAndSolveAgent(t *testing.T) {
	t.Run("plan", func(t *testing.T) {
		resp, err := plan("一个水果店周一卖出了15个苹果。周二卖出的苹果数量是周一的两倍。周三卖出的数量比周二少了5个。请问这三天总共卖出了多少个苹果？")
		assert.NoError(t, err)
		fmt.Println(resp)
	})
	t.Run("think", func(t *testing.T) {
		p := NewPlanAndSolveAgent("一个水果店周一卖出了15个苹果。周二卖出的苹果数量是周一的两倍。周三卖出的数量比周二少了5个。请问这三天总共卖出了多少个苹果？")
		p.Run()
	})
}
