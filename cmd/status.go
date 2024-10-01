package cmd

import (
	"cachprax/cmd/internal/state"
	"fmt"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the status of the caching proxy server",
	Long:  "Check if the caching proxy server is running and display its origin and port if it is.",
	RunE:  statusCommand,
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

func statusCommand(_ *cobra.Command, _ []string) error {
	serverInfo, err := state.GetDataFromFile()
	if err != nil {
		fmt.Println("Server status: not running (could not find server info state)")
		return nil
	}

	ok := state.IsProcessRunning(serverInfo.PID)
	if !ok {
		fmt.Println("Server status: not running")
		return nil
	}

	fmt.Println("Server status: running")
	fmt.Printf("Origin: %s\n", serverInfo.Origin)
	fmt.Printf("Port: %d\n", serverInfo.Port)

	return nil
}
