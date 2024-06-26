package config

import "github.com/spf13/viper"

type Config struct {
	DBHost string `mapstructure:"POSTGRES_HOST"`
	DBPort string `mapstructure:"POSTGRES_PORT"`
	DBUser string `mapstructure:"POSTGRES_USER"`
	DBPass string `mapstructure:"POSTGRES_PASSWORD"`
	DBName string `mapstructure:"POSTGRES_DB"`
	Port   string `mapstructure:"PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return config, nil
}
