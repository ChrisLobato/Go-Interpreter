package main

import (
	"fmt"
	"jade/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	//Welcome message
	fmt.Printf("Hello %s! This is the Jade Programming Language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)
}
