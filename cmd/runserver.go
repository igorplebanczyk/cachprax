package cmd

import (
	"cachprax/internal/cache"
	"cachprax/internal/server"
	"github.com/urfave/cli/v2"
	"net/url"
	"time"
)

func runserverCommand(c *cli.Context) error {
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
