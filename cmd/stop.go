package cmd

import (
	"cachprax/cmd/internal/file"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func stopCommand(_ *cobra.Command, _ []string) error {
	serverInfo, err := file.GetDataFromFile()
	if err != nil {
		return fmt.Errorf("server is not running or server info file is missing")
	}

	ok := file.IsProcessRunning(serverInfo.PID)
	if !ok {
		return fmt.Errorf("server is not running")
	}

	process, err := os.FindProcess(serverInfo.PID)
	if err != nil {
		return fmt.Errorf("could not find process: %v", err)
	}

	err = process.Kill()
	if err != nil {
		return fmt.Errorf("could not kill process: %v", err)
	}

	// Remove the PID file
	pidFile := filepath.Join(os.TempDir(), "cachprax.json")
	err = os.Remove(pidFile)
	if err != nil {
		return fmt.Errorf("could not remove JSON file: %v", err)
	}

	fmt.Print("Server stopped.")
	return nil
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the running caching proxy server",
	Long:  "Stop the running caching proxy server by killing the process and removing its PID file.",
	RunE:  stopCommand,
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
