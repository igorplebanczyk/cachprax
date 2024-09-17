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
		return fmt.Errorf("could not get server info from file: %v", err)
	}

	process, err := os.FindProcess(serverInfo.PID)
	if err != nil {
		return fmt.Errorf("could not find process: %v", err)
	}

	err = process.Kill()
	if err != nil {
		return fmt.Errorf("could not kill process: %v", err)
	}

	pidFile := filepath.Join(os.TempDir(), "cachprax.json")
	err = os.Remove(pidFile)
	if err != nil {
		return fmt.Errorf("could not remove JSON file: %v", err)
	}

	fmt.Println("Server stopped.")
	return nil
}
