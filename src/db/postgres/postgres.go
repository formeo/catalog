package postgres

import (
	"catalog/src/config"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBConnections struct {
	Master *sqlx.DB
}

func NewDBConnections(conf *config.PostgresConfig) (*DBConnections, error) {
	if conf.MasterDSN == "" {
		return nil, fmt.Errorf("empty master postgres DSN")
	}
	masterConn, err := connectDB(conf.MasterDSN, conf.MaxIdleConns, conf.MaxConns, conf.ConnMaxLifetime)
	if err != nil {
		return nil, err
	}

	return &DBConnections{
		Master: masterConn,
	}, nil
}

func connectDB(dsn string, maxIdleConns, maxConns int, maxLifetime time.Duration) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxConns)
	db.SetConnMaxLifetime(maxLifetime)

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
