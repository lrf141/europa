package main

import (
        "fmt"
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

func initConfig() *Config {

        buf, err := ioutil.ReadFile(confPath)
        if err != nil {
                panic(err.Error())
        }

        var conf *Config
        conf = new(Config)
        err = yaml.Unmarshal(buf, conf)
        if err != nil {
                panic(err.Error())
        }
        return conf
}

func (conf *Config) getDSN() string {
        return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
                conf.User,
                conf.Pass,
                conf.Host,
                conf.Port,
                conf.Db,
                )
}
