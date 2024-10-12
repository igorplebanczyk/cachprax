package cli

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "cachprax",
	Short: "A simple caching proxy",
	Long:  "A simple proxy server that forwards requests to a target server and caches the responses.",
}

func Execute() error {
	return rootCmd.Execute()
}
