package main

import (
	"cachprax/cmd"
	"cachprax/internal/config"
)

func main() {
	config.Init()
	_ = cmd.Execute()
}
