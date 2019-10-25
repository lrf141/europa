package main

import (
	"fmt"
	"github.com/urfave/cli"
	"time"
)

const migrateDir = "./migrations/migrate"

func migrateRunAction(c *cli.Context) {

	db := prepareDbDriver()
	fmt.Println(db.Driver)

	defer func() {
		err := db.Driver.Close()
		if err != nil {
			panic(err.Error())
		}
	}()

}

func migrateCreateAction(c *cli.Context) error {

	if c.NumFlags() < 1 || fileName == "" {
		err := cli.ShowCommandHelp(c, "create:migrate")
		if err != nil {
			panic(err.Error())
		}
		return cli.NewExitError("Please set migrations name.", 1)
	}

	if !isDirExist(migrateDir) {
		err := mkDir(migrateDir)
		if err != nil {
			panic(err.Error())
		}
	}

	t := time.Now().Format("20060102150405")
	migrateFile := fmt.Sprintf("%s/%s_%s", migrateDir, t, fileName)

	err := touchFile(migrateFile + upSql)
	if err != nil {
		panic(err.Error())
	}

	err = touchFile(migrateFile + downSql)
	if err != nil {
		err2 := deleteFile(migrateFile + upSql)
		if err2 != nil {
			panic(err2.Error())
		}
		panic(err.Error())
	}

	return nil
}
