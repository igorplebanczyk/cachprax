package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"net/http"
)

func cacheCommand(c *cli.Context) error {
	if !c.Bool("clear") {
		return nil
	}

	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:3001/cache/clear", nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return cli.Exit("Failed to clear the cache", 1)
	}

	fmt.Printf("Cache cleared\n")

	return nil
}
