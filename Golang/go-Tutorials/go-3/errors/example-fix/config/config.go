package config

import (
	"io/ioutil"
)

func Load() (string, error) {
	data, err := ioutil.ReadFile("/tmp/myHotelApp/config.txt")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
