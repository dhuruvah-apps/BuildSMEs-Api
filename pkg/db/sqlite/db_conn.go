package sqlite

import (
	"io/fs"
	"os"
	"path"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/dhuruvah-apps/BuildSMEs-Api/config"
)

const (
	maxOpenConns    = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

// Return new Postgresql db instance
func NewSqliteDB(c *config.Config) (*sqlx.DB, error) {
	err := os.MkdirAll(c.Sqlite.DbFilePath, fs.ModeDir)

	if err != nil {
		return nil, err
	}

	dataSourceName := path.Join(c.Sqlite.DbFilePath, c.Sqlite.DbFileName)

	db, err := sqlx.Connect(c.Sqlite.SqliteDriver, dataSourceName)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
