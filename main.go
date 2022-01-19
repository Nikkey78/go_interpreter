package main

import (
	"console"
	"fmt"
	"interpreter/repl"
	"os"
	"os/signal"
	"os/user"
)

func main() {
	defer console.GreenReset()()
	go stopApp()

	user, err := user.Current()
	if err != nil {
		console.ColorPrintln(console.Color_Red, err)
		return
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Println("Feel free to type in commands")

	repl.Start(os.Stdin, os.Stdout)
}

func stopApp() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	console.ResetColor()
	os.Exit(0)
}
