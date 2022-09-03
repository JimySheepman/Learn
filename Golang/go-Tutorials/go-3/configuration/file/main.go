package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Configuration struct {
	Server   Server `json:"server"`
	Database DB     `json:"database"`
}
type Server struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}
type DB struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	confFile, err := os.Open("myConf.json")
	if err != nil {
		panic(err)
	}
	defer confFile.Close()
	conf, err := ioutil.ReadAll(confFile)
	if err != nil {
		panic(err)
	}

	myConf := Configuration{}
	err = json.Unmarshal(conf, &myConf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", myConf)
}
