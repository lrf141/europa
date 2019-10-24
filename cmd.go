package main

import "github.com/urfave/cli"

func registerCommands(app *cli.App) {
        app.Commands = []cli.Command{
                {
                        Name: "run",
                        Usage: "Run Database Migrations",
                        Action: func(c *cli.Context) {},
                },
                {
                        Name: "rollback",
                        Usage: "Rollback Migrations",
                        Action: func(c *cli.Context) {},
                },
                {
                        Name: "status",
                        Usage: "Show Migrations Status",
                        Action: func(c *cli.Context) {},
                },
        }
}