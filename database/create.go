package database

import (
	"database/sql"
	"log"
	// Import postgres driver.
	_ "github.com/lib/pq"
)

func CreatePhrase(db *sql.DB, entry PhraseEntry) {
	const insertSQL =  `
INSERT INTO wubalubadubdub.words (phrase, file, startPhrase, endPhrase, nextPhrase) VALUES ($1, $2, $3, $4, $5);
`
	_, err := db.Exec(insertSQL, entry.Phrase, entry.File, entry.Start, entry.End, entry.Next)
	if err != nil {
		log.Println(err)
	}
}
