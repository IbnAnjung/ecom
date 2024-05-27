package config

type Config struct {
	Http       HttpConfig
	Mongo      MongoConfig
	Mysql      MysqlConfig
	ServiceURI ServiceBaseURI
}

type HttpConfig struct {
	Port string
}

type MongoConfig struct {
	Host              string
	User              string
	Password          string
	Source            string
	ProductCollection string
}

type MysqlConfig struct {
	User               string
	Password           string
	Host               string
	Port               string
	Schema             string
	Timeout            int
	MaxIddleConnection int
	MaxOpenConnection  int
	MaxLifeTime        int
}

type ServiceBaseURI struct {
	Authentication string
	Store          string
}
