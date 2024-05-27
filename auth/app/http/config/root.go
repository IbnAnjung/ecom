package config

import (
	"github.com/spf13/viper"
)

func LoadConfig() Config {
	viper.AutomaticEnv()
	return Config{
		Http:  loadHttpConfig(),
		Mysql: loadMysqlConfig(),
		Jwt:   loadJwtConfig(),
	}
}

func loadHttpConfig() HttpConfig {
	return HttpConfig{
		Port: viper.GetString("HTTP_PORT"),
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

func loadJwtConfig() JwtConfig {
	return JwtConfig{
		SecretKey:            viper.GetString("JWT_SECRET"),
		SellerSecretKey:      viper.GetString("JWT_SELLER_SECRET"),
		AccessTokenLifeTime:  viper.GetInt("JWT_ACCESS_TOKEN_LIFETIME"),
		RefreshTokenLifeTime: viper.GetInt("JWT_REFRESH_TOKEN_LIFETIME"),
	}
}
