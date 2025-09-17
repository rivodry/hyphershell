package main

import (
	"fmt"
	"hyphershell/command"
	"hyphershell/interpreter"
	"os"
	"strings"
)

var CMD command.Command

func shell() {
	for {

		fmt.Print(">")
		CMD.Get()

		interpreter.Interpret(CMD.Cmd)
	}
}
func e(args []string) {
	interpreter.Interpret(strings.Join(args, " "))
}
func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		shell()
	} else {
		if args[0] == "-e" {
			e(args[1:])
		}
	}
}
