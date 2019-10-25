package main

import (
	"fmt"
	"github.com/urfave/cli"
)

func registerCommands(app *cli.App) {
	app.Commands = []cli.Command{
		{
			Name:   "run:migrate",
			Usage:  "Run Database Migrations",
			Action: runActionHandler,
		},
		{
			Name:   "rollback:migrate",
			Usage:  "Rollback Migrations",
			Action: func(c *cli.Context) {},
		},
		{
			Name: "create:migrate",
			Usage: "Create Migrate SQL File into migrations/migrate",
			Action: func(c *cli.Context) {},
		},
		{
			Name: "run:seeds",
			Usage: "Run Database Seeds",
			Action: func(c *cli.Context) {},
		},
		{
			Name: "rollback:seeds",
			Usage: "Rollback Seeds",
			Action: func(c *cli.Context) {},
		},
		{
			Name: "create:seeds",
			Usage: "Create Seeds SQL File into migrations/seed",
			Action: func(c *cli.Context) {},
		},
		{
			Name:   "status",
			Usage:  "Show Migrations Status",
			Action: func(c *cli.Context) {},
		},
	}
}

func prepareDbDriver() *DB {

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

	db := prepareDbDriver()
	fmt.Println(db.Driver)

	defer func() {
		err := db.Driver.Close()
		if err != nil {
			panic(err.Error())
		}
	}()
}
