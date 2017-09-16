package main

import (
	"fmt"
	"github.com/sjfricke/wubalubadubdub/ingestion"
	"os"
	"sync"
)

func main() {
	w := ingestion.GetWatson()
	var wg sync.WaitGroup
	wg.Add(len(os.Args) - 1)
	for _, file := range os.Args[1:] {
		go func(f string) {
			defer wg.Done()
			tt := ingestion.GetWordLocations(w, f)
			for _, w := range tt.Words {
				fmt.Printf("%v%v\n", f, w)
			}
		}(file)
	}

	wg.Wait()
}
