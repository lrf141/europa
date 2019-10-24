package main

import "github.com/urfave/cli"

func registerCmdInfo(app *cli.App) {
        app.Name = "europa"
        app.Usage = "MySQL CLI Migration Tools"
        app.Author = "lrf141"
        app.Version = "1.0.0"
}
