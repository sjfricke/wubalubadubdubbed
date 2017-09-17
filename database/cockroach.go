package database

import (
	"database/sql"
	"log"
	"time"

	// Import postgres driver.
	_ "github.com/lib/pq"
)

type PhraseEntry struct {
	Phrase, File string
	Previous, Start, End, Next time.Time
}

func ConnectCockroach(dbURL string) *sql.DB {

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	return db;
}
