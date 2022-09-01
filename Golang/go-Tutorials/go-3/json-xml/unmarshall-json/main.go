package main

import (
	"encoding/json"
	"fmt"
)

type MyJson struct {
	Cat
}

type Cat struct {
	Name string
	Age  uint
}

func main() {
	myJson := []byte(`{"cat":{ "name":"Joey", "age":8}}`)
	c := MyJson{}
	err := json.Unmarshal(myJson, &c)
	if err != nil {
		panic(err)
	}
	fmt.Println(c.Cat.Name)
	fmt.Println(c.Cat.Age)
}
