package cmd

import (
	"cachprax/internal/cache"
	"cachprax/internal/server"
	"fmt"
	"github.com/urfave/cli/v2"
	"net/url"
	"time"
)

func runserverCommand(c *cli.Context) error {
	if !c.Bool("override") {
		return fmt.Errorf("\nthis command should not be ran manually; use the start command instead\n")
	}

	port := c.Int("port")
	originURL, err := url.Parse(c.String("origin"))
	if err != nil {
		return err
	}

	cfg := &server.Config{
		Port:   port,
		Origin: originURL,
		Cache:  cache.NewCache(5*time.Minute, 10*time.Minute),
	}

	return cfg.StartServer()
}
