package config

import "github.com/spf13/viper"

func GetEnvConfig() *viper.Viper {
	config := viper.New()
	config.SetConfigFile(".env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return config
}
