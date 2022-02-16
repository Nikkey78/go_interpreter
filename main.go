package main

import (
	"console"
	"fmt"
	"interpreter/repl"
	"os"
	"os/user"
)

func main() {
	defer console.GreenReset()()
	go console.StopApp()

	user, err := user.Current()
	if err != nil {
		console.ColorPrintln(console.Color_Red, err)
		return
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Println("Feel free to type in commands")

	repl.Start(os.Stdin, os.Stdout)
}
