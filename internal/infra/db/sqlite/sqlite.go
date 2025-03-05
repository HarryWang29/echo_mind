package sqlite

import (
	"fmt"

	"github.com/HarryWang29/sqlite"
	_ "github.com/mutecomm/go-sqlcipher"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SQLite struct {
	key    string
	path   string
	dbName string
	tx     *gorm.DB
}

func NewSQLite(key, path, dbName string) (s *SQLite, err error) {
	s = &SQLite{
		key:    key,
		path:   path,
		dbName: dbName,
	}
	dsn := fmt.Sprintf("%s/%s?_pragma_key=x'%s'", s.path, s.dbName, s.key)
	s.tx, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}
	return s, nil
}

func (s *SQLite) DB() *gorm.DB {
	return s.tx
}
