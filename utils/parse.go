package utils

import "regexp"

func NewRegexp(match string) *regexp.Regexp {
	return regexp.MustCompile(match)
}
