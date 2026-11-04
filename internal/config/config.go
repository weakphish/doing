package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() {
	// Name of the config file without an extension (Viper will intuit the type
	// from an extension on the actual file)
	viper.SetConfigName("config")

	// Add search paths to find the file
	viper.AddConfigPath("$HOME/.config/doing")
	viper.AddConfigPath(".")

	// Find and read the config file
	err := viper.ReadInConfig()

	// Handle errors
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
