package main

import (
	"fmt"
	"github.com/sjfricke/wubalubadubdub/database"
//	"github.com/sjfricke/wubalubadubdub/encoding"
	"time"
)

func main() {
	fmt.Println("hello")

	database.Test();

	// 2.418 seconds
	start := time.Date(2009, time.November, 10, 0, 0, 2, 418000, time.UTC)
	// 0.453
	end := time.Date(2009, time.November, 10, 0, 0, 0, 453000, time.UTC)

	next := time.Date(2009, time.November, 10, 0, 0, 0, 600000, time.UTC)

	db := database.ConnectCockroach("postgresql://root@FrickeFresh-Linux:26257?sslmode=disable");

//	e := new(database.PhraseEntry)

	entry := database.PhraseEntry{
		Phrase: "hello world",
		File: "/home/fricke/Videos/rm1.mp4",
		Start: start,
		End: end,
		Next: next,}

	
/*	
	entry := database.PhraseEntry()
	entry.phrase = "hello world";
	entry.file = "/home/fricke/Videos/rm1.mp4";
	entry.start =  start;
	entry.end = end;
	entry.next =next;
	database.CreatePhrase(db, entry)
*/

	database.CreatePhrase(db, entry)
		
	
//	encoding.Crop("/home/fricke/Videos/rm1.mp4", "/home/fricke/Videos/test.mp4", start, end);
}
