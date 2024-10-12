package cli

import (
	"cachprax/cli/internal/state"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the running caching proxy server",
	Long:  "Stop the running caching proxy server by killing the process and removing its PID state.",
	RunE:  stopCommand,
}

func init() {
	rootCmd.AddCommand(stopCmd)
}

func stopCommand(_ *cobra.Command, _ []string) error {
	serverInfo, err := state.GetDataFromFile()
	if err != nil {
		return fmt.Errorf("server is not running or server info state is missing")
	}

	ok := state.IsProcessRunning(serverInfo.PID)
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

	// Remove the PID state
	pidFile := filepath.Join(os.TempDir(), "cachprax.json")
	err = os.Remove(pidFile)
	if err != nil {
		return fmt.Errorf("could not remove JSON state: %v", err)
	}

	fmt.Printf("Server stopped.\n")
	return nil
}
