package agents

import "testing"

func TestReflectionAgent_Run(t *testing.T) {
	NewReflectionAgent("编写一个golang函数，找出1到n之间所有的素数", 2).Run()
}
