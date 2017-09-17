package database

import (
	"database/sql"
	"log"
	"strings"
//	"time"
	"fmt"

	// Import postgres driver.
	_ "github.com/lib/pq"
)

func ReadPhrase(db *sql.DB, phr string) PhraseEntry {

	phr = strings.Replace(phr, ";", "", -1)
	phr = strings.Replace(phr, "'", "", -1)
	
	rows, err := db.Query(fmt.Sprintf(`SELECT phrase, file, startPhrase, endPhrase, nextPhrase
                                           FROM wubalubadubdub.words
                                           WHERE (phrase LIKE '%s')
                                           ORDER BY RANDOM() LIMIT 20;`,
                                           strings.ToLower(strings.TrimSpace(phr))))

	if  err != nil {
		log.Println(err)
	}

	// hold a temp entry and return entry
	entry := PhraseEntry{}

	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&entry.Phrase, &entry.File, &entry.Start, &entry.End, &entry.Next); err != nil {
			log.Println(err)
		}

		// Currently just returning first foudn item
		break
	}

	err = rows.Err() // get encountered errors

	return entry
}
