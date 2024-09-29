package cmd

import (
	"cachprax/internal/cache"
	"cachprax/internal/server"
	"fmt"
	"github.com/spf13/cobra"
	"net/url"
	"time"
)

func runserverCommand(cmd *cobra.Command, _ []string) error {
	override, err := cmd.Flags().GetBool("override")
	if err != nil {
		return err
	}
	if !override {
		return fmt.Errorf("this command should not be run manually; use the start command instead")
	}

	port, err := cmd.Flags().GetInt("port")
	if err != nil {
		return err
	}

	origin, err := cmd.Flags().GetString("origin")
	if err != nil {
		return err
	}
	originURL, err := url.Parse(origin)
	if err != nil {
		return err
	}

	cacheExpire, err := cmd.Flags().GetInt("cache-expire")
	if err != nil {
		return err
	}
	cachePurge, err := cmd.Flags().GetInt("cache-purge")
	if err != nil {
		return err
	}

	cfg := &server.Config{
		Port:   port,
		Origin: originURL,
		Cache:  cache.NewCache(time.Duration(cacheExpire)*time.Minute, time.Duration(cachePurge)*time.Minute),
	}

	return cfg.StartServer()
}

var runserverCmd = &cobra.Command{
	Use:   "runserver",
	Short: "Start the proxy server manually (internal use)",
	Long:  "This command starts the proxy server manually, but it should be run only via the start command.",
	RunE:  runserverCommand,
}

func init() {
	rootCmd.AddCommand(runserverCmd)

	runserverCmd.Flags().Bool("override", false, "Override manual check (for internal use only)")
	runserverCmd.Flags().IntP("port", "p", 8080, "Port to run the server on")
	runserverCmd.Flags().StringP("origin", "o", "", "Origin server URL")
	runserverCmd.Flags().Int("cache-expire", 10, "Cache expiration duration in minutes")
	runserverCmd.Flags().Int("cache-purge", 30, "Cache purge interval in minutes")
}
