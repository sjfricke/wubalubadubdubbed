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

func ReadPhrase(db *sql.DB, phrases ...string) (PhraseEntry, int) {

	phr := phrases[0]
	
	phr = strings.Replace(phr, ";", "", -1)
	phr = strings.Replace(phr, "'", "", -1)
	
	rows, err := db.Query(fmt.Sprintf(`SELECT phrase, file, startPhrase, endPhrase, nextPhrase
                                           FROM wubalubadubdub.words
                                           WHERE (phrase LIKE '%s') AND (file LIKE '/home/fricke/videos/wubalubadubdub/s2e1.mp4')
                                           ORDER BY RANDOM() limit 15;`, // 15 to high or low?
		strings.ToLower(strings.TrimSpace(phr))))

	if  err != nil {
		log.Println(err)
	}

	// hold a temp entry and return entry
	entry := PhraseEntry{}
	cEntry := PhraseEntry{}
	var count int = 1
	var found bool = false
	
	defer rows.Close()
	for rows.Next() {
		
		if err = rows.Scan(&entry.Phrase, &entry.File, &entry.Start, &entry.End, &entry.Next); err != nil {
			log.Println(err)
		}

		cRows, cErr := db.Query(`SELECT phrase, file, startPhrase, endPhrase, nextPhrase
                                           FROM wubalubadubdub.words
                                           WHERE (file LIKE $1) AND (startPhrase > $2)
                                           ORDER BY startPhrase limit 3;`, entry.File, entry.Start)

		if  cErr != nil {
			log.Println(cErr)
		}

		// last items won't have chilren
		if len(phrases) <= count {
			log.Println("234324")
			break
		}  
		
		defer rows.Close()
		for cRows.Next() {
			if cErr = cRows.Scan(&cEntry.Phrase, &cEntry.File, &cEntry.Start, &cEntry.End, &cEntry.Next); cErr != nil {
				log.Println(cErr)
			}

			// hard check for 2 or 3 words in row
			// settle with anything we get more then one right away
			if (cEntry.Phrase == phrases[count] && (cEntry.Start).Equal(entry.Next)) {
				entry.Next = cEntry.Next
				count = count + 1
				found = true
				
				entry.Phrase = fmt.Sprintf("%s %s", entry.Phrase, cEntry.Phrase)
				// last items wonx't have chilren
				if len(phrases) <= count {
					break
				}  
			} else {
				break
			}
		}

		
		// will actaully never let a 3rd item in a row come now
		if found == true {
			
			log.Println("111")
			break
		}

	}
	
	err = rows.Err() // get encountered errors
	
	return entry, count
}
