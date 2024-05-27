package config

import (
	"github.com/spf13/viper"
)

func LoadConfig() Config {
	viper.AutomaticEnv()
	return Config{
		Mysql: loadMysqlConfig(),
	}
}

func loadMysqlConfig() MysqlConfig {
	return MysqlConfig{
		User:               viper.GetString("DB_USER"),
		Password:           viper.GetString("DB_PASSWORD"),
		Host:               viper.GetString("DB_HOST"),
		Schema:             viper.GetString("DB_SCHEMA"),
		Port:               viper.GetString("DB_PORT"),
		Timeout:            viper.GetInt("DB_TIMEOUT"),
		MaxIddleConnection: viper.GetInt("DB_MAX_IDDLE_CONNECTION"),
		MaxOpenConnection:  viper.GetInt("DB_MAX_OPEN_CONNECTION"),
		MaxLifeTime:        viper.GetInt("DB_MAX_LIFETIME"),
	}
}
