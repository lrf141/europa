package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	Conf   *Config
	Driver *sql.DB
}

const _driver = "mysql"

func initDb() (*DB, error) {

	conf := initConfig()

	driver, err := sql.Open(_driver, conf.getDSN())
	if err != nil {
		return nil, err
	}

	return &DB{
		Conf:   conf,
		Driver: driver,
	}, nil
}

func (db *DB) HealthCheck() error {
	return db.Driver.Ping()
}
