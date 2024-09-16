package cmd

import (
	"cachprax/internal/cache"
	"cachprax/internal/server"
	"github.com/urfave/cli/v2"
	"time"
)

func startCommand(c *cli.Context) error {
	cfg := &server.Config{
		Port:   c.Int("port"),
		Origin: c.String("origin"),
		Cache:  cache.NewCache(5*time.Minute, 10*time.Minute),
	}
	return cfg.StartServer()
}
