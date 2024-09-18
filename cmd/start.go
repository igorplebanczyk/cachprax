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
	_, err := file.GetDataFromFile()
	if err == nil {
		return fmt.Errorf("server already running")
	}

	port := c.Int("port")
	origin := c.String("origin")
	cacheExpire := c.Int("cache-expire")
	cachePurge := c.Int("cache-purge")

	// Get the path to the currently running binary
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("could not get executable path: %v", err)
	}

	// Start a new process for the server in the background
	cmd := exec.Command(exePath, "runserver", "--origin", origin, "--port", strconv.Itoa(port), "--override", "--cache-expire", strconv.Itoa(cacheExpire), "--cache-purge", strconv.Itoa(cachePurge))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return fmt.Errorf("could not start server process: %v", err)
	}

	fmt.Printf("Server started in background with PID: %d", cmd.Process.Pid)

	err = file.SaveDataToFile(cmd.Process.Pid, origin, port)
	if err != nil {
		return fmt.Errorf("could not save server info to file: %v", err)
	}

	return nil
}
