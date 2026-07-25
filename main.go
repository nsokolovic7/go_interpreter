package main

import (
	"fmt"
	"go_interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is my very own programming language!\n", user.Username)
	fmt.Printf("Type in a comamnd!\n")
	repl.Start(os.Stdin, os.Stdout)
}
