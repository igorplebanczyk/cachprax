package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
	"strconv"
)

func stopCommand(c *cli.Context) error {
	pidFile := filepath.Join(os.TempDir(), "cachprax.pid")

	// Read the PID from the file
	pidBytes, err := os.ReadFile(pidFile)
	if err != nil {
		return fmt.Errorf("could not read PID file: %v", err)
	}

	pid, err := strconv.Atoi(string(pidBytes))
	if err != nil {
		return fmt.Errorf("could not convert PID to integer: %v", err)
	}

	// Find the process and terminate it
	process, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("could not find process: %v", err)
	}

	err = process.Kill()
	if err != nil {
		return fmt.Errorf("could not kill process: %v", err)
	}

	err = os.Remove(pidFile)
	if err != nil {
		return fmt.Errorf("could not remove PID file: %v", err)
	}

	fmt.Println("Server stopped.")
	return nil
}
