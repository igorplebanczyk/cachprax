package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

func conntestCommand(cmd *cobra.Command, _ []string) error {
	origin, err := cmd.Flags().GetString("origin")
	if err != nil {
		return fmt.Errorf("error getting origin flag: %v", err)
	}

	fmt.Printf("Attempting to connect to %s...\n", origin)

	client := http.Client{}
	req, err := http.NewRequest(http.MethodHead, origin, nil)
	if err != nil {
		return fmt.Errorf("error creating a request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending a HEAD request: %v", err)
	}

	fmt.Printf("Sending a HEAD request...\n")

	if resp.StatusCode > 299 {
		return fmt.Errorf("error: %v", resp.Status)
	}

	fmt.Printf("Received status code %d\nConnection successful\n", resp.StatusCode)

	return nil
}

var conntest = &cobra.Command{
	Use:   "conntest",
	Short: "Test connection to the origin server",
	Long:  "Test connection to the origin server by sending a HEAD request",
	RunE:  conntestCommand,
}

func init() {
	rootCmd.AddCommand(conntest)

	conntest.Flags().StringP("origin", "o", "", "Origin server to test connection (required)")
	if err := conntest.MarkFlagRequired("origin"); err != nil {
		return
	}
}
