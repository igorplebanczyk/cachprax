package cmd

import (
	"cachprax/internal/cache"
	"cachprax/internal/server"
	"fmt"
	"github.com/urfave/cli/v2"
	"net/url"
	"time"
)

func startCommand(c *cli.Context) error {
	originURL, err := url.Parse(c.String("origin"))
	if err != nil {
		return err
	}

	cfg := &server.Config{
		Port:   c.Int("port"),
		Origin: originURL,
		Cache:  cache.NewCache(5*time.Minute, 10*time.Minute),
	}

	//go func() {
	//
	//}()

	err = cfg.StartServer()
	if err != nil {
		fmt.Printf("error starting server\n")
	}

	return nil
}
