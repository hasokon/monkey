package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/hasokon/monkey/repl"
)

func main() {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programinglanguage!\n", currentUser.Name)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
