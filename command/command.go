package command

import (
	"bufio"
	"os"
)

type Command struct {
	Cmd string
}

func (cmd *Command) Get() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	cmd.Cmd = input

}
