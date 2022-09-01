package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// username:password@protocol(address)/dbname?param=value
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/school")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.Open("create_table.sql")
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := db.Exec(string(b))
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = res
}
