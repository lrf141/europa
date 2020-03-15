package main

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"os"
)

const (
	upSql   = ".up.sql"
	downSql = ".down.sql"
)

const (
	messageFormat = "%s %s: %s %s"
)

func isDirExist(dirName string) bool {
	_, err := os.Stat(dirName)
	if err != nil {
		return false
	}
	return true
}

func mkDir(name string) error {
	return os.Mkdir(name, 0777)
}

func isFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if err != nil {
		return false
	}
	return true
}

func touchFile(name string) error {

	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	_, err = fmt.Fprintln(file, "")
	return err
}

func deleteFile(name string) error {
	return os.Remove(name)
}

func getFileNameWithoutExtension(target string, ext string) string {
	return target[:len(target)-len(ext)]
}

func skipMigrate(migrateName string) bool {
	return fileName != "" && fileName != migrateName
}

func printSuccessStatus(file string, types string, action string) {
	msg := fmt.Sprintf(messageFormat, action, types, file, aurora.Green("[Success]").String())
	fmt.Println(msg)
}

func printRollbackStatus(file string, types string, action string) {
	msg := fmt.Sprintf(messageFormat, action, types, file, aurora.Yellow("[Rollback]").String())
	fmt.Println(msg)
}

func printRollbackFailedStatus(file string, types string, action string) {
	msg := fmt.Sprintf(messageFormat, action, types, file, aurora.Red("[Rollback Failed]").String())
	fmt.Println(msg)
}

func printFailedStatus(file string, types string, action string) {
	msg := fmt.Sprintf(messageFormat, action, types, file, aurora.Red("[Failed]").String())
	fmt.Println(msg)
}

func printSkipStatus(file string, types string, action string) {
	msg := fmt.Sprintf(messageFormat, action, types, file, aurora.Blue("[Skip]").String())
	fmt.Println(msg)
}