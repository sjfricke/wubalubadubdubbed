package main

import (
	"github.com/sjfricke/wubalubadubdub/database"
//	"github.com/sjfricke/wubalubadubdub/encoding"
	"time"
)

func main() {

	db := database.ConnectCockroach("postgresql://root@FrickeFresh-Linux:26257?sslmode=disable");

	// 2.418 seconds
	start := time.Date(2009, time.November, 10, 0, 0, 2, 418000, time.UTC)
	// 0.453
	end := time.Date(2009, time.November, 10, 0, 0, 0, 453000, time.UTC)
	next := time.Date(2009, time.November, 10, 0, 0, 0, 719000, time.UTC)

	entry := database.PhraseEntry{
		Phrase: "hey rick",
		File: "/home/fricke/Videos/rm1.mp4",
		Start: start,
		End: end,
		Next: next,}

	database.CreatePhrase(db, entry)

	// 2nd entry

	// 2.418 seconds
	start = time.Date(2009, time.November, 10, 0, 0, 6, 910000, time.UTC)
	// 0.453
	end = time.Date(2009, time.November, 10, 0, 0, 0, 526000, time.UTC)
	next = time.Date(2009, time.November, 10, 0, 0, 0, 598000, time.UTC)

	entry = database.PhraseEntry{
		Phrase: "whatever",
		File: "/home/fricke/Videos/rm1.mp4",
		Start: start,
		End: end,
		Next: next,}

	database.CreatePhrase(db, entry)

//	encoding.Crop("/home/fricke/Videos/rm1.mp4", "/home/fricke/Videos/test.mp4", start, end);
}
