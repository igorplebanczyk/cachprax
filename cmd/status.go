package cmd

import (
	"cachprax/cmd/internal/file"
	"fmt"

	"github.com/spf13/cobra"
)

func statusCommand(cmd *cobra.Command, _ []string) error {
	serverInfo, err := file.GetDataFromFile()
	if err != nil {
		fmt.Println("Server status: not running (could not find server info file)")
		return nil
	}

	ok := file.IsProcessRunning(serverInfo.PID)
	if !ok {
		fmt.Println("Server status: not running")
		return nil
	}

	fmt.Println("Server status: running")
	fmt.Printf("Origin: %s\n", serverInfo.Origin)
	fmt.Printf("Port: %d\n", serverInfo.Port)

	return nil
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the status of the caching proxy server",
	Long:  "Check if the caching proxy server is running and display its origin and port if it is.",
	RunE:  statusCommand,
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
