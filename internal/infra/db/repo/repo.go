package repo

import (
	"errors"
	"github.com/HarryWang29/echo_mind/config"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewRepo(cfg *config.DataSourceConfig) (*gorm.DB, error) {
	switch cfg.Driver {
	case "mysql":
		return NewMysql(cfg.DSN)
	}
	return nil, errors.New("not support driver")
}

func NewMysql(dsn string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}

func GetOptions() []gen.DOOption {
	return []gen.DOOption{}
}
