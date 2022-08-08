package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Omar-Belghaouti/ez/repl"
)

func Run() {
	// command to run the interpreter on a file
	// e.g. ez file.ez
	file := flag.String("file", "", "path to file to run")
	flag.Parse()

	if *file != "" {
		if strings.HasSuffix(*file, ".ez") {
			if err := repl.RunFile(*file); err != nil {
				fmt.Printf("Error running file: %s\n", err)
				os.Exit(1)
			}
		} else {
			fmt.Println("File must have .ez extension")
			os.Exit(1)
		}
	} else {
		repl.Start(os.Stdin, os.Stdout)
	}
}
