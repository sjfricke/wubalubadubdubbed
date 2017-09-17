package encoding

import (
	"fmt"
	"time"
	"os"
	"path/filepath"
	"strconv"
	"github.com/sjfricke/wubalubadubdub/database"
)

//var startOffset time.Duration = (0 * time.Millisecond)

// timeDiff is used to find a differnce in time and return as a time value
func timeDiff(start time.Time, end time.Time) time.Time {
	delta := end.Sub(start)
	return time.Date(2009, time.November, 10, 0, 0, 0, 0, time.UTC).Add(delta)
}

// Encode takes slice of parsed entries and makes a single video
func Encode(entries []database.PhraseEntry) {
	ePath := filepath.Join(".", strconv.FormatInt(time.Now().Unix(), 10));
	os.MkdirAll(ePath, os.ModePerm)

	var cFiles []string
	var cPath string
	
	for i := 0; i < len(entries); i++ {
		cPath = fmt.Sprintf("./%s/%s",ePath, fmt.Sprintf("%d.mp4", i))
		fmt.Printf("DEBUG: cPath: %s\n", cPath);
		cFiles = append(cFiles, cPath)
		
		Crop(entries[i].File,
			cPath,
			entries[i].Start,
			timeDiff(entries[i].Start, entries[i].Next))
	}

	Stitch(cFiles, filepath.Join(ePath, "output.mp4"))	
}
