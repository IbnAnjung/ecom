package orm

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type IGormUow interface {
	Recovery(ctx context.Context)
	Begin(ctx context.Context) error
	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
	GetDB() *gorm.DB
}

type gormUow struct {
	db *gorm.DB
	tx *gorm.DB
}

type OrmDialect string

const (
	MysqlDialect OrmDialect = "mysql"
)

type GormConfig struct {
	Connection gorm.ConnPool
	Dialect    OrmDialect
}

func NewGormOrm(conf GormConfig) (orm IGormUow, err error) {
	var dialect gorm.Dialector

	if conf.Dialect == MysqlDialect {
		dialect = mysql.New(mysql.Config{
			Conn: conf.Connection,
		})
	} else {
		err = fmt.Errorf("unsupport orm dialect")
		return
	}

	g, err := gorm.Open(dialect, &gorm.Config{
		SkipDefaultTransaction: true,
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             1 * time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			}),
	})
	if err != nil {
		return
	}

	return &gormUow{
		db: g,
	}, nil
}

func (uow *gormUow) Recovery(ctx context.Context) {
	if r := recover(); r != nil {
		log.Printf("Gorm Recovery : %v", r)
		if uow.tx != nil {
			uow.tx.Rollback()
		}
		return
	}
}

func (uow *gormUow) Begin(ctx context.Context) error {
	uow.tx = uow.db.WithContext(ctx).Begin()
	return uow.tx.Error
}

func (uow *gormUow) Rollback(ctx context.Context) (err error) {
	err = uow.tx.WithContext(ctx).Rollback().Error
	uow.tx = nil
	return
}

func (uow *gormUow) Commit(ctx context.Context) (err error) {
	err = uow.tx.WithContext(ctx).Commit().Error
	uow.tx = nil
	return
}

func (uow *gormUow) GetDB() (db *gorm.DB) {
	db = uow.db
	if uow.tx != nil {
		db = uow.tx
	}

	return
}
