package main

import (
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/school")
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := db.Exec("UPDATE teacher SET firstname = ? WHERE id = ?", "Daniel", 1)
	if err != nil {
		fmt.Println(err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}

	if affected != 1 {
		fmt.Printf("Something went wrong %d rows were affected expected 1\n", affected)
	} else {
		fmt.Println("Update is a success")
	}

}
