package main

import (
        "fmt"
        "os"
)

const (
        upSql = ".up.sql"
        downSql = ".down.sql"
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
                panic(err.Error())
        }

        defer func() {
                err := file.Close()
                if err != nil {
                        panic(err.Error())
                }
        }()

        _, err = fmt.Fprintln(file, "")
        return err
}

func deleteFile(name string) error {
        return os.Remove(name)
}