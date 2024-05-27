package sql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

type SqlDriver string

type MysqlConfig struct {
	User               string
	Password           string
	Host               string
	Port               string
	DbName             string
	Loc                *time.Location
	Timeout            time.Duration
	MaxIddleConnection int
	MaxIddleTime       int
	MaxOpenConnection  int
	MaxLifeTime        int
}

type MysqlConnection struct {
	Db *sql.DB
}

func NewMysqlConnection(ctx context.Context, conf MysqlConfig) (c MysqlConnection, err error) {
	sqlConf := mysql.Config{
		User:      conf.User,
		Passwd:    conf.Password,
		DBName:    conf.DbName,
		Addr:      fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		Loc:       conf.Loc,
		Timeout:   conf.Timeout,
		ParseTime: true,
	}

	connector, err := mysql.NewConnector(&sqlConf)
	if err != nil {
		fmt.Printf("fail create new mysql connector: %s\n", sqlConf.Addr)
		return
	}

	db := sql.OpenDB(connector)

	if err = db.PingContext(ctx); err != nil {
		fmt.Printf("fail ping mysql connection: %s\n", sqlConf.Addr)
		return
	}

	db.SetConnMaxIdleTime(time.Duration(conf.MaxIddleTime))
	db.SetMaxIdleConns(conf.MaxIddleConnection)
	db.SetConnMaxLifetime(time.Duration(conf.MaxLifeTime))
	db.SetMaxOpenConns(conf.MaxOpenConnection)

	c.Db = db

	return
}

func (c *MysqlConnection) Cleanup() error {
	return c.Db.Close()
}
