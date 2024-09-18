package cmd

import (
	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	return &cli.App{
		Name:  "cachprax",
		Usage: "simple caching proxy server",
		Commands: []*cli.Command{
			{
				Name:   "start",
				Usage:  "Start the caching proxy server",
				Action: startCommand,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "origin",
						Usage:    "The URL of the origin server",
						Required: true,
					},
					&cli.IntFlag{
						Name:     "port",
						Usage:    "The port on which the caching proxy server will listen on",
						Required: true,
					},
					&cli.IntFlag{
						Name:        "cache-expire",
						Usage:       "The time in minutes after which the cache will expire; default is 5 minutes",
						DefaultText: "5",
						Required:    false,
					},
					&cli.IntFlag{
						Name:        "cache-purge",
						Usage:       "The time in minutes after which the cache will be purged; default is 10 minutes",
						DefaultText: "10",
						Required:    false,
					},
				},
			},
			{
				Name:   "stop",
				Usage:  "Stop the caching proxy server",
				Action: stopCommand,
			},
			{
				Name:   "status",
				Usage:  "Get the status of the caching proxy server",
				Action: statusCommand,
			},
			{
				Name:   "cache",
				Usage:  "Manage the cache",
				Action: cacheCommand,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:     "clear",
						Usage:    "Clear the cache",
						Required: false,
					},
					&cli.BoolFlag{
						Name:     "count",
						Usage:    "Get the number of items in the cache",
						Required: false,
					},
				},
			},
			{
				Name:   "conntest",
				Usage:  "Test the connection to the origin server",
				Action: conntestCommand,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "origin",
						Usage:    "The URL of the origin server",
						Required: true,
					},
				},
			},
			{
				Name:   "runserver",
				Usage:  "Run the caching proxy server - for internal use only",
				Action: runserverCommand,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "origin",
						Usage:    "The URL of the origin server",
						Required: true,
					},
					&cli.IntFlag{
						Name:     "port",
						Usage:    "The port on which the caching proxy server will listen on",
						Required: true,
					},
					&cli.IntFlag{
						Name:     "cache-expire",
						Usage:    "The time in seconds after which the cache will expire; default is 5 minutes",
						Required: false,
					},
					&cli.IntFlag{
						Name:     "cache-purge",
						Usage:    "The time in seconds after which the cache will be purged; default is 10 minutes",
						Required: false,
					},
					&cli.BoolFlag{
						Name:     "override",
						Usage:    "This command should not be ran manually; use this to override this behavior",
						Required: true,
					},
				},
			},
		},
	}
}
