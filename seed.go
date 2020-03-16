package main

import "github.com/urfave/cli"

const seedDir = "./migrations/seed"

type Seed struct {
        Db *DB
        Type string
        Action string
        Dir string
}

func initSeed(action string) (*Seed,error) {

        db := prepareDbDriver()
        if !isDirExist(seedDir) {
                return nil, cli.NewExitError("Does not exist "+seedDir, 1)
        }

        return &Seed{
                Db: db,
                Type: "seed",
                Action: action,
                Dir: seedDir,
        }, nil
}


func (m *Seed) Register(name string, flag int) {
        m.Db.RegisterSeed(name, flag)
}

func (m *Seed) Update(name string, flag int) {
        m.Db.UpdateSeedInfo(name, flag)
}

func (m *Seed) Create() error {
        return m.Db.CreateSeedSchema()
}

func (m *Seed) GetDb() *DB {
        return m.Db
}

func (m *Seed) GetRegister() (map[string]int, error) {
        return m.Db.GetRegisterSeeds()
}

func (m *Seed) GetType() string {
        return m.Type
}

func (m *Seed) GetAction() string {
        return m.Action
}

func (m *Seed) GetDir() string {
        return m.Dir
}

func (m *Seed) CloseDbDriver() error {
        return m.Db.Driver.Close()
}