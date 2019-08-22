package quick_tui

import (
	"fmt"
	"strconv"
	"strings"
)

func InputNumberResolution(axis string) int {
	for {
		fmt.Printf("# <Please enter the new %s, then press return> ", axis)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		val, err := strconv.Atoi(text)
		if err == nil && val > 0 {
			return val
		}
	}
}
