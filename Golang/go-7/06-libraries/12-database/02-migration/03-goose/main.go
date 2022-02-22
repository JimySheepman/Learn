package main

import (
	"fmt"
	"log"
	"os"

	_ "goose-sample/migrations"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

func main() {
	var command string
	fmt.Print("command (status|up|down): ")
	fmt.Scanf("%s", &command)

	db, err := goose.OpenDBWithDriver("sqlite3", os.TempDir()+"goose.db")
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	if err := goose.Run(command, db, "./migrations"); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
