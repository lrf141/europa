package main

import "github.com/urfave/cli"

const migrateDir = "./migrations/migrate"

type Migrate struct {
        Db *DB
        Type string
        Action string
        Dir string
}

func initMigrate(action string) (*Migrate,error) {

        db := prepareDbDriver()
        if !isDirExist(migrateDir) {
                return nil, cli.NewExitError("Does not exist "+migrateDir, 1)
        }

        return &Migrate{
                Db: db,
                Type: "migrate",
                Action: action,
                Dir: migrateDir,
        }, nil
}


func (m *Migrate) Register(name string, flag int) {
        m.Db.RegisterMigrate(name, flag)
}

func (m *Migrate) Update(name string, flag int) {
        m.Db.UpdateMigrateInfo(name, flag)
}

func (m *Migrate) Create() error {
        return m.Db.CreateMigrateSchema()
}

func (m *Migrate) GetDb() *DB {
        return m.Db
}

func (m *Migrate) GetRegister() (map[string]int, error) {
        return m.Db.GetRegisterMigrates()
}

func (m *Migrate) GetType() string {
        return m.Type
}

func (m *Migrate) GetAction() string {
        return m.Action
}

func (m *Migrate) GetDir() string {
        return m.Dir
}

func (m *Migrate) CloseDbDriver() error {
        return m.Db.Driver.Close()
}