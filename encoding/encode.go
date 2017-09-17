package encoding

import (
	"fmt"
	"time"
	"os"
	"path/filepath"
	"strconv"
	"github.com/sjfricke/wubalubadubdub/database"
)

func init() {
}


// Encode takes slice of parsed entries and makes a single video
func Encode(entries []database.PhraseEntry) {
	ePath := filepath.Join(".", strconv.FormatInt(time.Now().Unix(), 10))
	os.MkdirAll(ePath, os.ModePerm)

	var cFiles []string
	var cPath string
	
	for i := 0; i < len(entries); i++ {
		cPath = filepath.Join(ePath, fmt.Sprintf("%d.mp4", i))
		cFiles = append(cFiles, cPath)
		
		Crop(entries[i].File,
			cPath,
			entries[i].Start,
			entries[i].Next)
	}

	Stitch(cFiles, filepath.Join(ePath, "output.mp4"))	
}
