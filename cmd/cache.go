package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"net/http"
)

func cacheCommand(c *cli.Context) error {
	if c.Bool("clear") {
		err := clearCache()
		if err != nil {
			return err
		}
		return nil
	}

	if c.Bool("count") {
		err := countCache()
		if err != nil {
			return err
		}
		return nil
	}

	return nil
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

	fmt.Printf("Cache count: %v", bodyString)
	return nil
}
