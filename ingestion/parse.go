package ingestion

import (
	// "fmt"
	"github.com/cg505/watson-go-sdk"
	"time"
	"github.com/sjfricke/wubalubadubdub/database"
	"database/sql"
)

func ParseAndAdd(w *watson.Watson, file string, db *sql.DB) {
	tt := GetWordLocations(w, file)
	for i, w := range tt.Words {
		if(w.Confidence >= 0.8) {
			start := time.Date(2009, time.November, 10, 0, 0, 0, 0, time.UTC).Add(time.Duration(w.Begin * 1000) * time.Millisecond)
			end := time.Date(2009, time.November, 10, 0, 0, 0, 0, time.UTC).Add(time.Duration(w.End * 1000) * time.Millisecond)
			var next time.Time
			if (i < len(tt.Words) - 1) {
				next = time.Date(2009, time.November, 10, 0, 0, 0, 0, time.UTC).Add(time.Duration(tt.Words[i + 1].Begin * 1000) * time.Millisecond)
			} else {
				next = end
			}
			entry := database.PhraseEntry{
				Phrase: w.Token,
				File: file,
				Start: start,
				End: end,
				Next: next,
			}
			database.CreatePhrase(db, entry)
		}
	}
}
