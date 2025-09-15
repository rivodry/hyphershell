package main

import (
	"fmt"
)

func shell() {
	for {
		var cmd string
		fmt.Print(">")
		fmt.Scanln(&cmd)
		fmt.Println(cmd)
	}
}
func main() {
	shell()

}
