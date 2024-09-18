package main

import (
	"cachprax/cmd"
	"fmt"
	"os"
)

func main() {
	cliApp := cmd.NewApp()
	err := cliApp.Run(os.Args)
	if err != nil {
		fmt.Printf("error running cli: %v", err)
	}

	_, err = fmt.Fprintf(os.Stdout, "\n") // Add a line break for cleaner output
	if err != nil {
		return
	}
}
