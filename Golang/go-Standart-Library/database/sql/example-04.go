package main

import (
	"database/sql"
	"log"
)

func main() {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
	}
	id := 37
	_, execErr := tx.Exec(`UPDATE users SET status = ? WHERE id = ?`, "paid", id)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}
