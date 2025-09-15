package main

import (
	"fmt"
	"hyphershell/command"
	"hyphershell/interpreter"
)

var CMD command.Command

func shell() {
	for {

		fmt.Print(">")
		CMD.Get()
		interpreter.Interpret(CMD.Cmd)
	}
}
func main() {
	shell()

}
