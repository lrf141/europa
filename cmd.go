package main

import (
	"fmt"
	"github.com/urfave/cli"
)

func registerCommands(app *cli.App) {
	app.Commands = []cli.Command{
		{
			Name:   "run",
			Usage:  "Run Database Migrations",
			Action: runActionHandler,
		},
		{
			Name:   "rollback",
			Usage:  "Rollback Migrations",
			Action: func(c *cli.Context) {},
		},
		{
			Name:   "status",
			Usage:  "Show Migrations Status",
			Action: func(c *cli.Context) {},
		},
	}
}

func rootActionHandler() *DB {

	db, err := initDb()

	if err != nil {
		panic(err.Error())
	}

	err = db.HealthCheck()
	if err != nil {
		driverErr := db.Driver.Close()
		if driverErr != nil {
			panic(driverErr.Error())
		}
		panic(err.Error())
	}

	return db
}

func runActionHandler(c *cli.Context) {

	db := rootActionHandler()
	fmt.Println(db.Driver)

	defer func() {
		err := db.Driver.Close()
		if err != nil {
			panic(err.Error())
		}
	}()
}
