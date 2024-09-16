package cmd

import "github.com/urfave/cli/v2"

func NewApp() *cli.App {
	return &cli.App{
		Name:  "cachprax",
		Usage: "A simple caching proxy",
		Commands: []*cli.Command{
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
				},
			},
		},
		Flags: []cli.Flag{},
	}
}
