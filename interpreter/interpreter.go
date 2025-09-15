package interpreter

import "fmt"

func Interpret(cmd string) {

	fmt.Println(Lex(cmd))
}
