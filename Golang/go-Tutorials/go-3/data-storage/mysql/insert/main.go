package main

import "database/sql"

func createTeacher(firstname string, lastname string, db *sql.DB) (int64, error) {
	res, err := db.Exec("INSERT INTO `teacher` (`create_time`, `firstname`, `lastname`) VALUES (NOW(), ?, ?)", firstname, lastname)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
