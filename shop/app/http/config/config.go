package config

type Config struct {
	Http       HttpConfig
	Mysql      MysqlConfig
	ServiceURI ServiceBaseURI
}

type HttpConfig struct {
	Port string
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
}
