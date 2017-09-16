package database

import (
	"database/sql"
	"log"
	"strings"
	
	// Import postgres driver.
	_ "github.com/lib/pq"
)

const  _createSQL =  `
INSERT INTO wubalubadubdub.words (phrase, file, startPhrase, endPhrase, nextPhrase) VALUES ($1, $2, $3, $4, $5);
`

func CreatePhrase(db *sql.DB, entry PhraseEntry) {

	// Keeping all words lowercase
	// also trimming for good mesaures
	_, err := db.Exec( _createSQL,
		strings.ToLower(strings.TrimSpace(entry.Phrase)),
		strings.ToLower(strings.TrimSpace(entry.File)),
		entry.Start, entry.End, entry.Next)
	
	if err != nil {
		log.Println(err)
	}
}
