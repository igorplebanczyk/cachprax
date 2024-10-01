package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func Init() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not get user home directory: %v", err)
	}

	configPath := filepath.Join(home, "cachprax.yml")

	// Check if the config state exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Create the config state with default values
		defaultConfig := []byte(`default_port: 8080
								cache_port: 3001
								cache_expire: 10
								cache_purge: 30
								origin: ""
								`)
		err = os.WriteFile(configPath, defaultConfig, 0644)
		if err != nil {
			return fmt.Errorf("could not write default config state: %v", err)
		}
	}

	viper.SetConfigName("cachprax")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(home)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("could not read config file: %v", err)
	}

	return nil
}
