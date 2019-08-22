package quick_tui

import (
	"bufio"
	"fmt"
	"strings"
)

var reader *bufio.Reader

func SetReader(r *bufio.Reader) {
	reader = r
}

func InputYesOrNo() bool {
	for {
		fmt.Print("# <Type: y or n, then press return> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "y" {
			return true
		} else if text == "n" {
			return false
		}
	}
}
