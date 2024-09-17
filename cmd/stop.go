package cmd

import (
	"cachprax/cmd/internal/file"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
)

func stopCommand(c *cli.Context) error {
	serverInfo, err := file.GetDataFromFile()
	if err != nil {
		return fmt.Errorf("\ncould not get server info from file: %v\n", err)
	}

	process, err := os.FindProcess(serverInfo.PID)
	if err != nil {
		return fmt.Errorf("\ncould not find process: %v\n", err)
	}

	err = process.Kill()
	if err != nil {
		return fmt.Errorf("\ncould not kill process: %v\n", err)
	}

	pidFile := filepath.Join(os.TempDir(), "cachprax.json")
	err = os.Remove(pidFile)
	if err != nil {
		return fmt.Errorf("\ncould not remove JSON file: %v\n", err)
	}

	fmt.Println("Server stopped.")
	return nil
}
