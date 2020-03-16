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
const seedSchema = "seed_schema"

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

func (db *DB) CreateSchema(schema string, types string) error {
	_, err := db.Driver.Query("create table if not exists " + schema + "(" + types + " text, flag tinyint(1) default 0)")
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) CreateMigrateSchema() error {
	return db.CreateSchema(migrateSchema, "migrate")
}

func (db *DB) CreateSeedSchema() error {
	return db.CreateSchema(seedSchema, "seed")
}

func (db *DB) Exec(query string) error {
	_, err := db.Driver.Query(query)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) GetRegister(schema string, types string) (map[string]int, error) {
	result, err := db.Driver.Query("select " + types + ",flag from " + schema)
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

func (db *DB) GetRegisterMigrates() (map[string]int, error) {
	return db.GetRegister(migrateSchema, "migrate")
}

func (db *DB) GetRegisterSeeds() (map[string]int, error) {
	return db.GetRegister(seedSchema, "seed")
}

func (db *DB) Register(schema string, name string, flag int, types string) {
	_, err := db.Driver.Query("insert into "+schema+"(" + types + ", flag) values (?, ?)", name, flag)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (db *DB) RegisterMigrate(migrate string, flag int) {
	db.Register(migrateSchema, migrate, flag, "migrate")
}

func (db *DB) RegisterSeed(seed string, flag int) {
	db.Register(seedSchema, seed, flag, "seed")
}

func (db *DB) UpdateMigrateInfo(migrate string, flag int) {
	db.UpdateInfo(migrateSchema, migrate, flag)
}

func (db *DB) UpdateSeedInfo(seed string, flag int) {
	db.UpdateInfo(seedSchema, seed, flag)
}

func (db *DB) UpdateInfo(schema string, name string, flag int) {
	_, err := db.Driver.Query("update "+schema+" set flag=? where migrate = ?", flag, name)
	if err != nil {
		log.Fatal(err.Error())
	}
}
