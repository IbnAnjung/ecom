package config

type Config struct {
	Http  HttpConfig
	Mysql MysqlConfig
	Jwt   JwtConfig
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

type JwtConfig struct {
	SecretKey            string
	SellerSecretKey      string
	AccessTokenLifeTime  int
	RefreshTokenLifeTime int
}
