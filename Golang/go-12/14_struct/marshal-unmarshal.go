package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type person struct {
	First       string `json:"first"`
	Last        string `json:"last"`
	Age         int    `json:"age"`
	NotExported string `json:"code"`
}

type persons struct {
	First string
	Last  string
	Age   int
}

func main() {
	// * marshal
	p1 := person{"James", "Bond", 20, "007"}
	json.NewEncoder(os.Stdout).Encode(p1)
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))

	// * unmarshal
	var p2 persons
	rdr := strings.NewReader(`{"First":"James", "Last":"Bond", "Age":20}`)
	json.NewDecoder(rdr).Decode(&p2)
	bs = []byte(`{"First":"James", "Last":"Bond", "Age":20}`)
	json.Unmarshal(bs, &p2)
	fmt.Println(p2.First)
	fmt.Println(p2.Last)
	fmt.Println(p2.Age)
}
