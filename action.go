package main

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/urfave/cli"
	"io/ioutil"
	"strings"
	"time"
)

type Action interface {
	Register(name string, flag int)
	Update(name string, flag int)
	Create() error
	GetDb() *DB
	GetRegister() (map[string]int, error)
	GetType() string
	GetAction() string
	GetDir() string
	CloseDbDriver() error
}

func runAction(c *cli.Context, action Action) error {

	defer func() {
		err := action.CloseDbDriver()
		if err != nil {
			panic(err)
		}
	}()

	register, err := action.GetRegister()
	if err != nil {
		return err
	}

	files, err := ioutil.ReadDir(action.GetDir())
	if err != nil {
		return err
	}

	for _, file := range files {

		name := getFileNameWithoutExtension(file.Name(), upSql)

		if strings.HasSuffix(file.Name(), downSql) || skipMigrate(name) {
			continue
		}

		flag, ok := register[name]
		if flag == 1 {
			printSkipStatus(file.Name(), action.GetType(), action.GetAction())
			continue
		}

		query, err := ioutil.ReadFile(action.GetDir() + "/" + file.Name())
		if err != nil {
			printFailedStatus(file.Name(), action.GetType(), action.GetAction())
			fmt.Println(err)
			continue
		}

		err = action.GetDb().Exec(string(query))
		if err != nil {
			printFailedStatus(file.Name(), action.GetType(), action.GetAction())
			fmt.Println(err)
			continue
		}
		printSuccessStatus(file.Name(), action.GetType(), action.GetAction())

		if ok {
			action.Update(name, 1)
		} else {
			action.Register(name, 1)
		}

		if fileName != "" {
			break
		}
	}

	return nil
}

func migrateRunAction(c *cli.Context) error {
	migrateAction, err := initMigrate("Run")
	if err != nil {
		panic(err)
	}
	return runAction(c, migrateAction)
}


func seedRunAction(c *cli.Context) error {
	seedAction, err := initSeed("Run")
	if err != nil {
		return err
	}
	return runAction(c, seedAction)
}

func rollbackAction(c *cli.Context, action Action) error {

	defer func() {
		err := action.CloseDbDriver()
		if err != nil {
			panic(err)
		}
	}()

	files, err := ioutil.ReadDir(action.GetDir())
	if err != nil {
		panic(err)
	}

	migrates, err := action.GetRegister()
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
			printSkipStatus(file.Name(), action.GetType(), action.GetAction())
			continue
		}

		query, err := ioutil.ReadFile(action.GetDir() + "/" + file.Name())
		if err != nil {
			printFailedStatus(file.Name(), action.GetType(), action.GetAction())
			fmt.Println(err)
			continue
		}

		err = action.GetDb().Exec(string(query))
		if err != nil {
			printFailedStatus(file.Name(), action.GetType(), action.GetAction())
			continue
		}
		printSuccessStatus(file.Name(), action.GetType(), action.GetAction())

		if ok {
			action.Update(migrateName, 0)
		}

		if fileName != "" {
			break
		}
	}

	return nil
}

func migrateRollbackAction(c *cli.Context) error {
	migrateAction, err := initMigrate("Rollback")
	if err != nil {
		panic(err)
	}
	return rollbackAction(c, migrateAction)
}

func seedRollbackAction(c *cli.Context) error {
	seedAction, err := initSeed("Rollback")
	if err != nil {
		return err
	}
	return rollbackAction(c, seedAction)
}

func createAction(c *cli.Context, action Action) error {

	err := action.Create()
	if err != nil {
		return err
	}

	defer func() {
		err := action.CloseDbDriver()
		if err != nil {
			panic(err)
		}
	}()

	if c.NumFlags() < 1 || fileName == "" {
		err := cli.ShowCommandHelp(c, "create:" + action.GetType())
		if err != nil {
			panic(err.Error())
		}
		return cli.NewExitError("Please set migrations or seed name.", 1)
	}

	if !isDirExist(action.GetDir()) {
		err := mkDir(action.GetDir())
		if err != nil {
			return err
		}
	}

	t := time.Now().Format("20060102150405")
	migrateFile := fmt.Sprintf("%s/%s_%s", action.GetDir(), t, fileName)

	err = touchFile(migrateFile + upSql)
	if err != nil {
		printFailedStatus(migrateFile+upSql, action.GetType(), action.GetAction())
		panic(err.Error())
	}
	printSuccessStatus(migrateFile+upSql, action.GetType(), action.GetAction())

	err = touchFile(migrateFile + downSql)
	if err != nil {
		printFailedStatus(migrateFile+downSql, action.GetType(), action.GetAction())
		printRollbackStatus(migrateFile+upSql, action.GetType(), action.GetAction())
		err2 := deleteFile(migrateFile + upSql)
		if err2 != nil {
			printRollbackFailedStatus(migrateFile+upSql, action.GetType(), action.GetAction())
			return err2
		}
		return err
	}
	printSuccessStatus(migrateFile+downSql, action.GetType(), action.GetAction())

	return nil
}

func migrateCreateAction(c *cli.Context) error {
	migrateAction, err := initMigrate("Create")
	if err != nil {
		panic(err)
	}
	return createAction(c, migrateAction)
}


func seedCreateAction(c *cli.Context) error {
	seedAction, err := initSeed("Create")
	if err != nil {
		return err
	}
	return createAction(c, seedAction)
}

func statusAction(c *cli.Context, action Action) error {
	err := action.Create()
	if err != nil {
		return err
	}

	register, err := action.GetRegister()
	if err != nil {
		return err
	}

	for key, val := range register {
		var status string
		if val == 1 {
			status = aurora.Green("Active").String()
		} else {
			status = aurora.Yellow("Inactive").String()
		}

		fmt.Println(key + " " + status)
	}
	return nil
}

func migrateStatusAction(c *cli.Context) error {
	migrateAction, err := initMigrate("Status")
	if err != nil {
		return err
	}
	return statusAction(c, migrateAction)
}

func seedStatusAction(c *cli.Context) error {
	seedAction, err := initSeed("Status")
	if err != nil {
		return err
	}
	return statusAction(c, seedAction)
}
