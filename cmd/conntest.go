package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"net/http"
)

func conntestCommand(c *cli.Context) error {
	origin := c.String("origin")
	fmt.Printf("Attempting to connect to %s...\n", origin)

	client := http.Client{}
	req, err := http.NewRequest(http.MethodHead, origin, nil)
	if err != nil {
		return fmt.Errorf("\nerror creating a request: %v\n", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("\nerror sending a HEAD request: %v\n", err)
	}

	fmt.Printf("Sending a HEAD request...\n")

	if resp.StatusCode > 299 {
		return fmt.Errorf("\nerror: %v\n", resp.Status)
	}

	fmt.Printf("Received status code %d\nConnection successful\n", resp.StatusCode)

	return nil
}
