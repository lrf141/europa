package main

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/urfave/cli"
	"io/ioutil"
	"strings"
	"time"
)

const migrateDir = "./migrations/migrate"

func migrateRunAction(c *cli.Context) error {

	db := prepareDbDriver()

	defer func() {
		err := db.Driver.Close()
		if err != nil {
			panic(err.Error())
		}
	}()

	// if not exist table
	err := db.CreateMigrateSchema()
	if err != nil {
		panic(err.Error())
	}

	if !isDirExist(migrateDir) {
		return cli.NewExitError("Does not exist " + migrateDir, 1)
	}

	files, err := ioutil.ReadDir(migrateDir)
	if err != nil {
		panic(err.Error())
	}

	migrates, err := db.GetRegisterMigrates()

	for _, file := range files {

		if strings.HasSuffix(file.Name(), downSql) {
			continue
		}

		flag, ok := migrates[file.Name()]
		if flag == 1 {
			continue
		}

		query, err := ioutil.ReadFile(migrateDir+"/"+file.Name())
		if err != nil {
			panic(err.Error())
		}

		err = db.ExecMigrate(string(query))
		if err != nil {
			fmt.Println("Migrate " + file.Name() + aurora.Red("[Failed]").String())
			continue
		}
		fmt.Println("Migrate " + file.Name() + aurora.Green("[Success]").String())

		if ok {
			db.UpdateMigrateInfo(file.Name())
		} else {
			db.RegisterMigrate(file.Name(), 1)
		}
	}

	return nil
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
		fmt.Println("Create Migrate: "+ migrateFile + upSql + " " + aurora.Red("[Failed]").String())
		panic(err.Error())
	}
	fmt.Println("Create Migrate: "+ migrateFile + upSql + " " + aurora.Green("[Success]").String())

	err = touchFile(migrateFile + downSql)
	if err != nil {
		fmt.Println("Create Migrate: "+ migrateFile + downSql + " " + aurora.Red("[Failed]").String())
		fmt.Println("Create Migrate: "+ migrateFile + upSql + " " + aurora.Yellow("[Rollback]").String())
		err2 := deleteFile(migrateFile + upSql)
		if err2 != nil {
			fmt.Println("Create Migrate: "+ migrateFile + upSql + " " + aurora.Red("[Rollback Failed]").String())
			panic(err2.Error())
		}
		panic(err.Error())
	}
	fmt.Println("Create Migrate: "+ migrateFile + downSql + " " + aurora.Green("[Success]").String())

	return nil
}
