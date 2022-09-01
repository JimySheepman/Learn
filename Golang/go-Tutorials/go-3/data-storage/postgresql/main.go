package main

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
)

type Teacher struct {
	id        int
	firstname string
	lastname  string
}

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres dbname=school password=pg123 sslmode=disable")
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

	res, err = db.Exec("UPDATE teacher SET firstname = $1 WHERE id = $2", "Daniel", 1)

	affected, err := res.RowsAffected()
	_ = affected
}

func createTeacher(firstname string, lastname string, db *sql.DB) (int, error) {
	insertedId := 0
	err := db.QueryRow("INSERT INTO public.teacher (create_time, firstname, lastname) VALUES (NOW(),$1, $2) RETURNING id;", firstname, lastname).Scan(&insertedId)
	if err != nil {
		return 0, err
	}
	if insertedId == 0 {
		return 0, errors.New("something went wrong id inserted is equal to zero")
	}
	return insertedId, nil
}

func teacher(id int, db *sql.DB) (*Teacher, error) {
	teacher := Teacher{}
	err := db.QueryRow("SELECT id, firstname, lastname FROM teacher WHERE id > $1 ", id).Scan(&teacher.id, &teacher.firstname, &teacher.lastname)
	if err != nil {
		return &teacher, err
	}
	return &teacher, nil
}
