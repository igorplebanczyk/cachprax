package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var configFileName = "cachprax"
var configFileType = "yaml"
var defaultConfig = []byte(`proxy_port: 3000
cache_port: 3001
cache_expire: 10
cache_purge: 30
origin: ""
`)

func Init() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not get user home directory: %v", err)
	}

	configFile := configFileName + "." + configFileType
	configPath := filepath.Join(home, configFile)

	// Create default config file if it does not exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		err = os.WriteFile(configPath, defaultConfig, 0644)
		if err != nil {
			return fmt.Errorf("could not write default config state: %v", err)
		}
	}

	viper.SetConfigName(configFileName)
	viper.SetConfigType(configFileType)
	viper.AddConfigPath(home)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("could not read config file: %v", err)
	}

	return nil
}
