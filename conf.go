package main

import (
        "gopkg.in/yaml.v3"
        "io/ioutil"
)

type Config struct {
        Db string `yaml:"db"`
        Host string `yaml:"host"`
        Port int `yaml:"port"`
        User string `yaml:"user"`
        Pass string `yaml:"pass"`
}

const confPath = "./migrations/db.yaml"

var conf Config

func initConfig() {

        buf, err := ioutil.ReadFile(confPath)
        if err != nil {
                panic(err.Error())
        }

        err = yaml.Unmarshal(buf, &conf)
        if err != nil {
                panic(err.Error())
        }
}
