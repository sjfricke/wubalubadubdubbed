package ingestion

import (
	// "fmt"
	"log"
	"os"
	"github.com/cg505/watson-go-sdk"
	"github.com/sjfricke/wubalubadubdub/encoding"
	"strings"
)

func GetWatson() (*watson.Watson) {
	return watson.New(username, password)
}

func GetWordLocations(w *watson.Watson, file string) (*watson.Text) {
	basename := strings.Split(file, ".")[0]
	audiofile := strings.Join([]string{basename, "mp3"}, ".")
	encoding.Ffmpeg("-i", file, audiofile)

	is, err := os.Open(audiofile)
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
