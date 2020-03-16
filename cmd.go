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
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "name, n",
					Value: "",
					Usage: "Migration Name. ex) --name 20200101000000_create_user_table",
					Destination: &fileName,
				},
			},
		},
		{
			Name:   "rollback:migrate",
			Usage:  "Rollback Migrations",
			Action: migrateRollbackAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "name, n",
					Value: "",
					Usage: "Migration Name. ex) --name 20200101000000_create_user_table",
					Destination: &fileName,
				},
			},
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
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "name, n",
					Value: "",
					Usage: "Seed Name. ex) --name 20200101000000_create_user_table",
					Destination: &fileName,
				},
			},
		},
		{
			Name:   "rollback:seed",
			Usage:  "Rollback Seeds",
			Action: seedRollbackAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "name, n",
					Value: "",
					Usage: "Seed Name. ex) --name 20200101000000_create_user_table",
					Destination: &fileName,
				},
			},
		},
		{
			Name:   "create:seed",
			Usage:  "Create Seed SQL File into migrations/migrate",
			Action: seedCreateAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "name, n",
					Value:       "",
					Usage:       "Seed Name",
					Destination: &fileName,
				},
			},
		},
		{
			Name:   "status:migrate",
			Usage:  "Show Migrations Status",
			Action: migrateStatusAction,
		},
		{
			Name:   "status:seed",
			Usage:  "Show Seeds Status",
			Action: seedStatusAction,
		},
	}
}

func prepareDbDriver() *DB {

	db, err := initDb()

	if err != nil {
		panic(err)
	}

	err = db.HealthCheck()
	if err != nil {
		driverErr := db.Driver.Close()
		if driverErr != nil {
			panic(driverErr)
		}
		panic(err)
	}

	err = db.CreateMigrateSchema()
	if err != nil {
		panic(err)
	}

	return db
}
