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

func runAction(c *cli.Context, dir string, types string) error {
	action := "Run"
	if !isDirExist(dir) {
		return cli.NewExitError("Does not exist "+dir, 1)
	}

	db := prepareDbDriver()
	defer func() {
		err := db.Driver.Close()
		if err != nil {
			panic(err)
		}
	}()
	migrates, err := db.GetRegisterMigrates()
	if err != nil {
		panic(err)
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {

		migrateName := getFileNameWithoutExtension(file.Name(), upSql)

		if strings.HasSuffix(file.Name(), downSql) || skipMigrate(migrateName) {
			continue
		}

		flag, ok := migrates[migrateName]
		if flag == 1 {
			printSkipStatus(file.Name(), types, action)
			continue
		}

		query, err := ioutil.ReadFile(dir + "/" + file.Name())
		if err != nil {
			printFailedStatus(file.Name(), types, action)
			fmt.Println(err)
			continue
		}

		err = db.Exec(string(query))
		if err != nil {
			printFailedStatus(file.Name(), types, action)
			fmt.Println(err)
			continue
		}
		printSuccessStatus(file.Name(), types, action)

		if ok {
			db.UpdateMigrateInfo(migrateName, 1)
		} else {
			db.RegisterMigrate(migrateName, 1)
		}

		if fileName != "" {
			break
		}
	}

	return nil
}

func migrateRunAction(c *cli.Context) error {
	return runAction(c, migrateDir, "migrate")
}

func rollbackAction(c *cli.Context, dir string, types string) error {
	action := "Rollback"

	if !isDirExist(dir) {
		return cli.NewExitError("Does not exist "+dir, 1)
	}

	db := prepareDbDriver()

	defer func() {
		err := db.Driver.Close()
		if err != nil {
			panic(err)
		}
	}()

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	migrates, err := db.GetRegisterMigrates()
	if err != nil {
		panic(err)
	}

	for _, file := range files {

		migrateName := getFileNameWithoutExtension(file.Name(), downSql)

		if strings.HasSuffix(file.Name(), upSql) || skipMigrate(migrateName) {
			continue
		}

		flag, ok := migrates[migrateName]
		if flag == 0 || !ok {
			printSkipStatus(file.Name(), types, action)
			continue
		}

		query, err := ioutil.ReadFile(dir + "/" + file.Name())
		if err != nil {
			printFailedStatus(file.Name(), types, action)
			fmt.Println(err)
			continue
		}

		err = db.Exec(string(query))
		if err != nil {
			printFailedStatus(file.Name(), types, action)
			continue
		}
		printSuccessStatus(file.Name(), types, action)

		if ok {
			db.UpdateMigrateInfo(migrateName, 0)
		}

		if fileName != "" {
			break
		}
	}

	return nil
}

func migrateRollbackAction(c *cli.Context) error {
	return rollbackAction(c, migrateDir, "migrate")
}

func createAction(c *cli.Context, dir string, types string) error {
	action := "Create"

	if c.NumFlags() < 1 || fileName == "" {
		err := cli.ShowCommandHelp(c, "create:"+types)
		if err != nil {
			panic(err.Error())
		}
		return cli.NewExitError("Please set migrations name.", 1)
	}

	if !isDirExist(dir) {
		err := mkDir(dir)
		if err != nil {
			panic(err.Error())
		}
	}

	t := time.Now().Format("20060102150405")
	migrateFile := fmt.Sprintf("%s/%s_%s", dir, t, fileName)

	err := touchFile(migrateFile + upSql)
	if err != nil {
		printFailedStatus(migrateFile+upSql, types, action)
		panic(err.Error())
	}
	printSuccessStatus(migrateFile+upSql, types, action)

	err = touchFile(migrateFile + downSql)
	if err != nil {
		printFailedStatus(migrateFile+downSql, types, action)
		printRollbackStatus(migrateFile+upSql, types, action)
		err2 := deleteFile(migrateFile + upSql)
		if err2 != nil {
			printRollbackFailedStatus(migrateFile+upSql, types, action)
			panic(err2.Error())
		}
		panic(err.Error())
	}
	printSuccessStatus(migrateFile+downSql, types, action)

	return nil
}

func migrateCreateAction(c *cli.Context) error {
	return createAction(c, migrateDir, "migrate")
}

func migrateStatusAction(c *cli.Context) {

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

	migrates, err := db.GetRegisterMigrates()
	if err != nil {
		panic(err.Error())
	}

	for key, val := range migrates {

		var status string

		if val == 1 {
			status = aurora.Green("Active").String()
		} else {
			status = aurora.Yellow("Inactive").String()
		}

		fmt.Println(key + " " + status)
	}

}
