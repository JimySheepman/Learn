package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func IsValueInSlice(slice []int, value int) (result bool) {
	for _, n := range slice {
		if n == value {
			return true
		}

	}
	return false

}

type User struct {
	Id        int     `json:"id"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Balance   float64 `json:"balance"`
}

type AllUsers struct {
	Users []*User
}

func ShowAllUsers() (au *AllUsers) {
	file, err := os.OpenFile("list.json", os.O_RDWR|os.O_APPEND, 0666)
	checkError(err)
	b, err := ioutil.ReadAll(file)
	var alUsrs AllUsers
	json.Unmarshal(b, &alUsrs.Users)
	checkError(err)
	return &alUsrs
}
