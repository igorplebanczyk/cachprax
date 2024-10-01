package main

import (
	"cachprax/cmd"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func main() {
	initConfig()
	_ = cmd.Execute()
}

func initConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error finding home directory: %v\n", err)
		os.Exit(1)
	}

	configPath := filepath.Join(home, "cachprax.yml")

	// Check if the config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Create the config file with default values
		defaultConfig := []byte(`default_port: 8080
cache_port: 3001
cache_expire: 10
cache_purge: 30
origin: "http://example.com"
`)
		err = os.WriteFile(configPath, defaultConfig, 0644)
		if err != nil {
			fmt.Printf("Error creating config file: %v\n", err)
			os.Exit(1)
		}
	}

	viper.SetConfigName("cachprax")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(home)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		os.Exit(1)
	}
}
