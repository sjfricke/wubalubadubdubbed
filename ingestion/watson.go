package ingestion

import (
	// "fmt"
	"log"
	"os"
	"github.com/cg505/watson-go-sdk"
)

func GetWatson() (*watson.Watson) {
	return watson.New(username, password)
}

func GetWordLocations(w *watson.Watson, file string) (*watson.Text) {
	is, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer is.Close()

	tt, err := w.Recognize(is, "en-US_BroadbandModel", "mp3")
	if err != nil {
		log.Fatal(err)
	}

	return tt
}
