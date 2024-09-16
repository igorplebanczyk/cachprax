package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"net/http"
)

func ConntestCommand(c *cli.Context) error {
	origin := c.String("origin")
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
