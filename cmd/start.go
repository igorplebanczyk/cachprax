package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

func startCommand(c *cli.Context) error {
	port := c.Int("port")
	origin := c.String("origin")

	// Get the path to the currently running binary
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("could not get executable path: %v", err)
	}

	// Start a new process for the server in the background
	cmd := exec.Command(exePath, "runserver", "--origin", origin, "--port", strconv.Itoa(port))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return fmt.Errorf("could not start server process: %v", err)
	}

	fmt.Printf("Server started in background with PID: %d\n", cmd.Process.Pid)

	// Store the PID in a file for later use (e.g., for stop or status commands)
	pidFile := filepath.Join(os.TempDir(), "cachprax.pid")
	err = os.WriteFile(pidFile, []byte(fmt.Sprintf("%d", cmd.Process.Pid)), 0644)
	if err != nil {
		return fmt.Errorf("could not write PID file: %v", err)
	}

	return nil
}
