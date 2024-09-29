package main

import (
	"cachprax/cmd"
	"fmt"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
