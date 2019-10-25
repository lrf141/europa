package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {

	app := cli.NewApp()
	registerCmdInfo(app)
	registerCommands(app)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err.Error())
	}
}
