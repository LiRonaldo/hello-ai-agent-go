package constants

import (
	"fmt"
	"testing"
)

func Test_REACT_PROMPT_TEMPLATE(t *testing.T) {
	fmt.Println(fmt.Sprintf(ReactPromptTemplate, "1", "2", "3", "4", "5"))
}
