package main

import (
	"cachprax/cmd"
	"cachprax/internal/config"
	"fmt"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Printf("could not initialize config: %v\n", err)
		return
	}

	err = cmd.Execute()
	if err != nil {
		fmt.Printf("could not execute command: %v\n", err)
		return
	}
}
