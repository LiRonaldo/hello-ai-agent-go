package utils

import (
	"regexp"
	"strings"
)

func NewRegexp(match string) *regexp.Regexp {
	return regexp.MustCompile(match)
}

// 按中文句号自动换行（如需处理英文句号，替换为 "." 即可）
func WrapByPeriod(text string) string {
	// 替换所有中文句号为“句号+换行符”
	// 若要处理英文句号：strings.ReplaceAll(text, ".", ".\n")
	return strings.ReplaceAll(text, "。", "。\n")
}
