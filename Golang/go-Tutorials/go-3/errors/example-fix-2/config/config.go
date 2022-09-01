package config

import (
	"errors"
	"io/ioutil"
)

const fileHeader = "APPCONF"

func Load() (string, error) {
	data, err := ioutil.ReadFile("/tmp/myHotelApp/config.txt")
	if err != nil {
		return "", err
	}
	conf := string(data)
	if conf[0:8] != fileHeader {
		return "", errors.New("the config file do not begin by accepted header")
	}
	return conf, nil
}
