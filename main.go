package main

import (
	"fmt"
	"os/user"

	"github.com/Omar-Belghaouti/ez/cmd"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s to EZ programming language!\n", user.Username)
	fmt.Print("You can type in your commands\n")
	cmd.Run()
}
