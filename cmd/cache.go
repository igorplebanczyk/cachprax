package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

func cacheCommand(cmd *cobra.Command, _ []string) error {
	clearFlag, err := cmd.Flags().GetBool("clear")
	if err != nil {
		return err
	}

	countFlag, err := cmd.Flags().GetBool("count")
	if err != nil {
		return err
	}

	if clearFlag {
		err := clearCache()
		if err != nil {
			return err
		}
		return nil
	}

	if countFlag {
		err := countCache()
		if err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("no operation specified, please use --clear or --count")
}

func clearCache() error {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:3001/cache/clear", nil)
	if err != nil {
		return fmt.Errorf("\nfailed to create a request: %v\n", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("\nfailed to send a request: %v\n", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("\nfailed to clear the cache: %v\n", err)
	}

	fmt.Printf("Cache cleared\n")
	return nil
}

func countCache() error {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:3001/cache/count", nil)
	if err != nil {
		return fmt.Errorf("failed to create a request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send a request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get the cache count: %v", err)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read the response body: %v", err)
	}
	bodyString := string(bodyBytes)

	fmt.Printf("Cache count: %v\n", bodyString)
	return nil
}

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Manage cache operations",
	Long:  "Clear the cache or retrieve the cache item count from the caching proxy server.",
	RunE:  cacheCommand,
}

func init() {
	rootCmd.AddCommand(cacheCmd)

	cacheCmd.Flags().BoolP("clear", "c", false, "Clear the cache")
	cacheCmd.Flags().BoolP("count", "n", false, "Show the number of items in the cache")
}
