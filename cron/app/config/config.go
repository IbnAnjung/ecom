package config

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

type Config struct {
	Mysql MysqlConfig
}
