package cmd

import (
	"cachprax/internal/config"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration values",
	Long:  "Manage configuration values in the cachprax.yml file.",
	RunE:  configCommand,
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().BoolP("reset", "r", false, "Reset the configuration file to default values")
	configCmd.Flags().BoolP("set", "s", false, "Set a configuration value")
}

func configCommand(cmd *cobra.Command, args []string) error {
	reset, err := cmd.Flags().GetBool("reset")
	if err != nil {
		return fmt.Errorf("could not get reset flag: %v", err)
	}
	set, err := cmd.Flags().GetBool("set")
	if err != nil {
		return fmt.Errorf("could not get set flag: %v", err)
	}

	if reset {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("could not get user home directory: %v", err)
		}

		configFile := config.FileName + "." + config.FileType
		configPath := filepath.Join(home, configFile)

		err = os.WriteFile(configPath, config.DefaultValues, 0644)
		if err != nil {
			return fmt.Errorf("could not write default config file: %v", err)
		}

		fmt.Println("Configuration file reset to default values.")
		return nil
	}

	if set {
		if len(args) != 2 {
			return fmt.Errorf("set flag requires a key and a value")
		}

		key := args[0]
		value := args[1]

		viper.Set(key, value)

		err = viper.WriteConfig()
		if err != nil {
			return fmt.Errorf("could not write config file: %v", err)
		}

		fmt.Printf("Config value %s set to %s\n", key, value)
		return nil
	}

	return fmt.Errorf("either --reset or --set flag must be provided")
}
