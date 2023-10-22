package config

import "github.com/spf13/viper"

var config *viper.Viper

/// InitConfig configures sensible defaults for the web server
func InitConfig() {
	config = viper.New()

	// Host settings
	config.SetDefault("host.addr", "0.0.0.0")
	config.SetDefault("host.port", 3000)
}

func GetConfig() *viper.Viper {
	return config
}
