package command

import "fmt"

type Command struct {
	Cmd string
}

func (cmd *Command) Get() {
	fmt.Scanln(&cmd.Cmd)
}
