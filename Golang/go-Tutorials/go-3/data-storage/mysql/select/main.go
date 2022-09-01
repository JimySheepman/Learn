package main

import "database/sql"

type Teacher struct {
	id        int
	firstname string
	lastname  string
}

func teacher(id int, db *sql.DB) (*Teacher, error) {
	teacher := Teacher{id: id}
	err := db.QueryRow("SELECT firstname, lastname FROM teacher WHERE id = ?", id).Scan(&teacher.firstname, &teacher.lastname)
	if err != nil {
		return &Teacher{}, err
	}
	return &teacher, nil
}
