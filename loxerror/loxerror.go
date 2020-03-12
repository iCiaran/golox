package loxerror

import (
	"fmt"
)

var (
	HadError bool = false
)

func Error(line int, where, message string) {
	fmt.Printf("[%d] Error %s: %s\n", line, where, message)
	HadError = true
}
