package main

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/urfave/cli"
	"time"
)

const seedDir = "./migrations/seed"

func seedRunAction(c *cli.Context) {

	db := prepareDbDriver()
	fmt.Println(db.Driver)

	defer func() {
		err := db.Driver.Close()
		if err != nil {
			panic(err.Error())
		}
	}()
}

func seedCreateAction(c *cli.Context) error {

	if c.NumFlags() < 1 || fileName == "" {
		err := cli.ShowCommandHelp(c, "create:seed")
		if err != nil {
			panic(err.Error())
		}
		return cli.NewExitError("Please set seeds name.", 1)
	}

	if !isDirExist(seedDir) {
		err := mkDir(seedDir)
		if err != nil {
			panic(err.Error())
		}
	}

	t := time.Now().Format("20060102150405")
	seedFile := fmt.Sprintf("%s/%s_%s", seedDir, t, fileName)

	err := touchFile(seedFile + upSql)
	if err != nil {
		fmt.Println("Create Seed: "+ seedFile + upSql + " " + aurora.Red("[Failed]").String())
		panic(err.Error())
	}
	fmt.Println("Create Seed: "+ seedFile + upSql + " " + aurora.Green("[Success]").String())

	err = touchFile(seedFile + downSql)
	if err != nil {
		fmt.Println("Create Seed: "+ seedFile + downSql + " " + aurora.Red("[Failed]").String())
		fmt.Println("Create Seed: "+ seedFile + upSql + " " + aurora.Yellow("[Rollback]").String())
		err2 := deleteFile(seedFile + upSql)
		if err2 != nil {
			fmt.Println("Create Seed: "+ seedFile + upSql + " " + aurora.Red("[Rollback Failed]").String())
			panic(err2.Error())
		}
		panic(err.Error())
	}
	fmt.Println("Create Seed: "+ seedFile + downSql + " " + aurora.Green("[Success]").String())
	return nil
}
