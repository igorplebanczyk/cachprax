package cmd

import (
	"cachprax/cmd/internal/file"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
	"strconv"
)

func startCommand(c *cli.Context) error {
	port := c.Int("port")
	origin := c.String("origin")

	// Get the path to the currently running binary
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("\ncould not get executable path: %v\n", err)
	}

	// Start a new process for the server in the background
	cmd := exec.Command(exePath, "runserver", "--origin", origin, "--port", strconv.Itoa(port), "--override")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return fmt.Errorf("\ncould not start server process: %v\n", err)
	}

	fmt.Printf("Server started in background with PID: %d\n", cmd.Process.Pid)

	err = file.SaveDataToFile(cmd.Process.Pid, origin, port)
	if err != nil {
		return fmt.Errorf("\ncould not save server info to file: %v\n", err)
	}

	return nil
}
