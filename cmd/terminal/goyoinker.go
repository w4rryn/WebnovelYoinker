package main

import (
	"fmt"

	"github.com/lethal-bacon0/WebnovelYoinker/pkg/terminal"
	"github.com/lethal-bacon0/WebnovelYoinker/pkg/yoinker"
)

func main() {
	yoinker.MessageCallback = func(s string) {
		fmt.Println(s)
	}
	terminal.StartTerminal()
}
