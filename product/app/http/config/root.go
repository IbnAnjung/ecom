package config

import (
	"github.com/spf13/viper"
)

func LoadConfig() Config {
	viper.AutomaticEnv()
	return Config{
		Http:       loadHttpConfig(),
		Mongo:      loadMongoConfig(),
		Mysql:      loadMysqlConfig(),
		ServiceURI: loadServiceBaseURI(),
	}
}

func loadHttpConfig() HttpConfig {
	return HttpConfig{
		Port: viper.GetString("HTTP_PORT"),
	}
}

func loadMongoConfig() MongoConfig {
	return MongoConfig{
		Host:              viper.GetString("MONGO_HOST"),
		User:              viper.GetString("MONGO_USER"),
		Password:          viper.GetString("MONGO_PASSWORD"),
		Source:            viper.GetString("MONGO_SOURCE"),
		ProductCollection: viper.GetString("MONGO_COLLECTION_PRODUCT"),
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

func loadServiceBaseURI() ServiceBaseURI {
	return ServiceBaseURI{
		Authentication: viper.GetString("AUTH_SERVICE_BASE_URI"),
		Store:          viper.GetString("STORE_SERVICE_BASE_URI"),
	}
}
