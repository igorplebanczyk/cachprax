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
			},
		},
	}
}
