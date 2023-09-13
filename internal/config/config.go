package config

import "github.com/spf13/viper"

var config *viper.Viper

func InitConfig() {
	config = viper.New()
}

func GetConfig() *viper.Viper {
	return config
}
