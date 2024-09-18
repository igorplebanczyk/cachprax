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
		return fmt.Errorf("this command should not be ran manually; use the start command instead")
	}

	port := c.Int("port")
	originURL, err := url.Parse(c.String("origin"))
	if err != nil {
		return err
	}
	cacheExpire := time.Duration(c.Int("cache-expire")) * time.Minute
	cachePurge := time.Duration(c.Int("cache-purge")) * time.Minute

	cfg := &server.Config{
		Port:   port,
		Origin: originURL,
		Cache:  cache.NewCache(cacheExpire, cachePurge),
	}

	return cfg.StartServer()
}
