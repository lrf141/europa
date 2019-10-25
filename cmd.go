package main

import (
	"github.com/urfave/cli"
)

var (
	fileName string
)

func registerCommands(app *cli.App) {
	app.Commands = []cli.Command{
		{
			Name:   "run:migrate",
			Usage:  "Run Database Migrations",
			Action: migrateRunAction,
		},
		{
			Name:   "rollback:migrate",
			Usage:  "Rollback Migrations",
			Action: func(c *cli.Context) {},
		},
		{
			Name:   "create:migrate",
			Usage:  "Create Migrate SQL File into migrations/migrate",
			Action: migrateCreateAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "name, n",
					Value:       "",
					Usage:       "Migration Name",
					Destination: &fileName,
				},
			},
		},
		{
			Name:   "run:seed",
			Usage:  "Run Database Seeds",
			Action: seedRunAction,
		},
		{
			Name:   "rollback:seed",
			Usage:  "Rollback Seeds",
			Action: func(c *cli.Context) {},
		},
		{
			Name:   "create:seed",
			Usage:  "Create Seeds SQL File into migrations/seed",
			Action: seedCreateAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "name, n",
					Value:       "",
					Usage:       "Seeds Name",
					Destination: &fileName,
				},
			},
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
