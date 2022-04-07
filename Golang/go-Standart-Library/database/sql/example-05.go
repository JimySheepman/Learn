package main

import (
	"log"
)

func main() {
	id := 47
	result, err := db.ExecContext(ctx, "UPDATE balances SET balance = balance + 10 WHERE user_id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rows != 1 {
		log.Fatalf("expected to affect 1 row, affected %d", rows)
	}
}
