package main

import (
	"fmt"
	"os"
	"os/user"
	"monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Monkey Programming Language\nUser: %s\n", user.Username)
	fmt.Printf("use interactively below, Ctrl+D to exit\n")
	repl.Start(os.Stdin, os.Stdout)
}
