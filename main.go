package main

import (
        "github.com/urfave/cli"
        "os"
)

func main() {

        initConfig()

        app := cli.NewApp()
        registerCmdInfo(app)
        registerCommands(app)

        err := app.Run(os.Args)
        if err != nil {
                panic(err.Error())
        }
}