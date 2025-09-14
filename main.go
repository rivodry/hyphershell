package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func shell() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	cmd := make(chan string, 1)
	go func() {
		for {
			fmt.Print(">")
			var cm string
			fmt.Scan(&cm)
			cmd <- cm

		}
	}()
	for {
		select {
		case <-sigs:
			{
				os.Exit(0)
			}
		case cm := <-cmd:
			{
				if cm == "exit" {
					os.Exit(0)
				}
				fmt.Println(cm)
			}
		}
	}
}
func main() {
	shell()

}
