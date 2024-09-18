package cmd

import (
	"cachprax/cmd/internal/file"
	"fmt"
	"github.com/urfave/cli/v2"
)

func statusCommand(c *cli.Context) error {
	serverInfo, err := file.GetDataFromFile()
	ok := file.IsProcessRunning(serverInfo.PID)

	if err != nil || !ok {
		fmt.Printf("Server status: not running")
		return nil
	}

	fmt.Printf("Server status: running\n")
	fmt.Printf("Origin: %s\n", serverInfo.Origin)
	fmt.Printf("Port: %d\n", serverInfo.Port)

	return nil
}
