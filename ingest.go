package main

import (
	// "fmt"
	"github.com/sjfricke/wubalubadubdub/ingestion"
	"os"
	"sync"
	"github.com/sjfricke/wubalubadubdub/database"
)

func main() {
	w := ingestion.GetWatson()
	db := database.ConnectCockroach("postgresql://root@localhost:26257?sslmode=disable");
	var wg sync.WaitGroup
	wg.Add(len(os.Args) - 1)
	for _, file := range os.Args[1:] {
		go func(f string) {
			defer wg.Done()
			ingestion.ParseAndAdd(w, f, db)
		}(file)
	}

	wg.Wait()
}
