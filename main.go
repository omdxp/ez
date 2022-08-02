package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Omar-Belghaouti/ez/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s to EZ programming language!\n", user.Username)
	fmt.Print("You can type in your commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
