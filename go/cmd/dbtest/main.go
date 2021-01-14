package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("SQLite test program!")

	_ = os.Remove("./test.db")

	db, err := sqlx.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(1)

	_, err = db.Exec(`
		CREATE TABLE example (
			example_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			data TEXT NOT NULL
		)
	`)
	if err != nil {
		log.Fatalf("create table failed: %+v\n", err)
	}

	_, err = db.Exec(`
		INSERT INTO example (data)
		VALUES ('one'), ('two'), ('three')
	`)
	if err != nil {
		log.Fatalf("insert values failed: %+v\n", err)
	}

	var values []struct {
		Identifier int    `db:"example_id"`
		DataValue  string `db:"data"`
	}

	err = db.Select(&values, `
		SELECT
			example_id,
			data
		FROM example
	`)
	if err != nil {
		log.Fatalf("fetch values failed: %+v\n", err)
	}

	log.Printf("got values: %+v\n", values)
}
