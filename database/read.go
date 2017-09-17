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

	rows, err := db.Query(fmt.Sprintf(`SELECT phrase, file, startPhrase, endPhrase, nextPhrase
                                           FROM wubalubadubdub.words
                                           WHERE (phrase LIKE '%s');`,
                                           strings.ToLower(strings.TrimSpace(phr))))

	if  err != nil {
		log.Println(err)
	}

	// hold a temp entry and return entry
	entry := PhraseEntry{}
//	rEntry := PhraseEntry{}

	// used to hold the longest phrase so we can get the word with longest pattern
//	var lPhrase int = 0;

	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&entry.Phrase, &entry.File, &entry.Start, &entry.End, &entry.Next); err != nil {
			log.Println(err)
		}

		// Currently just returning first foudn item
		break

		// check if this row has longer span of words
//		count := len(strings.Split(entry.Phrase, " "))
//		if lPhrase < count {
//			lPhrase = count
//			rEntry.Phrase = entry.Phrase
//			rEntry.File = entry.File
//			rEntry.Start = entry.Start
//			rEntry.End = entry.End
//			rEntry.Next = entry.Next
//		}
	}

	err = rows.Err() // get encountered errors

	return entry
}
