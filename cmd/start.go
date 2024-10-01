package cmd

import (
	"cachprax/cmd/internal/file"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"strconv"
)

func startCommand(cmd *cobra.Command, _ []string) error {
	_, err := file.GetDataFromFile()
	if err == nil {
		return fmt.Errorf("server already running")
	}

	port, err := cmd.Flags().GetInt("port")
	if port == -1 {
		port = viper.GetInt("default_port")
	}
	if err != nil {
		return err
	}

	origin, err := cmd.Flags().GetString("origin")
	if origin == "" {
		origin = viper.GetString("origin")
	}
	if err != nil {
		return err
	}

	cacheExpire, err := cmd.Flags().GetInt("cache-expire")
	if cacheExpire == -1 {
		cacheExpire = viper.GetInt("cache_expire")
	}
	if err != nil {
		return err
	}

	cachePurge, err := cmd.Flags().GetInt("cache-purge")
	if cachePurge == -1 {
		cachePurge = viper.GetInt("cache_purge")
	}
	if err != nil {
		return err
	}

	// Get the path to the currently running binary
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("could not get executable path: %v", err)
	}

	// Start a new process for the server in the background
	runserverCmd := exec.Command(exePath, "runserver", "--origin", origin, "--port", strconv.Itoa(port), "--override", "--cache-expire", strconv.Itoa(cacheExpire), "--cache-purge", strconv.Itoa(cachePurge))
	runserverCmd.Stdout = os.Stdout
	runserverCmd.Stderr = os.Stderr

	err = runserverCmd.Start()
	if err != nil {
		return fmt.Errorf("could not start server process: %v", err)
	}

	fmt.Printf("Server started in background with PID: %d\n", runserverCmd.Process.Pid)

	err = file.SaveDataToFile(runserverCmd.Process.Pid, origin, port)
	if err != nil {
		return fmt.Errorf("could not save server info to file: %v", err)
	}

	return nil
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the caching proxy server",
	Long:  "Start the caching proxy server with the specified origin, port, cache expiration, and cache purge settings.",
	RunE:  startCommand,
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().StringP("origin", "o", "", "Origin server URL (required)")
	startCmd.Flags().IntP("port", "p", -1, "Port to run the server on")
	startCmd.Flags().Int("cache-expire", -1, "Cache expiration duration in minutes")
	startCmd.Flags().Int("cache-purge", -1, "Cache purge interval in minutes")
}
