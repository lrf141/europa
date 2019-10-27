package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DB struct {
	Conf   *Config
	Driver *sql.DB
}

const _driver = "mysql"
const migrateSchema = "migrate_schema"

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

func (db *DB) CreateMigrateSchema() error {
	_, err := db.Driver.Query("create table if not exists "+migrateSchema+"(migrate text, flag tinyint(1) default 0)")
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) ExecMigrate(query string) error {
	_, err := db.Driver.Query(query)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) GetRegisterMigrates() (map[string]int, error) {
	result, err := db.Driver.Query("select migrate,flag from "+migrateSchema)
	if err != nil {
		return nil, err
	}

	var migrates = map[string]int{}

	for result.Next() {
		var migrate string
		var flag int

		err := result.Scan(&migrate, &flag)
		if err != nil {
			log.Fatal(err.Error())
			continue
		}

		migrates[migrate] = flag
	}

	return migrates, nil
}

func (db *DB) RegisterMigrate(migrate string) {
	_, err := db.Driver.Query("insert into " + migrateSchema + "(migrate, flag) values (?, 1)", migrate)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (db *DB) UpdateMigrateInfo(migrate string) {
	_, err := db.Driver.Query("update "+ migrateSchema +" set flag=1 where migrate = ?", migrate)
	if err != nil {
		log.Fatal(err.Error())
	}
}