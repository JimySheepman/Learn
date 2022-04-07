package main

import (
	"database/sql"
	"log"
)

func main() {
	stmt, err := db.PrepareContext(ctx, "SELECT username FROM users WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	id := 43
	var username string
	err = stmt.QueryRowContext(ctx, id).Scan(&username)
	switch {
	case err == sql.ErrNoRows:
		log.Fatalf("no user with id %d", id)
	case err != nil:
		log.Fatal(err)
	default:
		log.Printf("username is %s\n", username)
	}
}
