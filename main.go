package main

import (
	"fmt"
	"hyphershell/command"
)

var CMD command.Command

func shell() {
	for {

		fmt.Print(">")
		CMD.Get()
		fmt.Println(CMD.Cmd)
	}
}
func main() {
	shell()

}
