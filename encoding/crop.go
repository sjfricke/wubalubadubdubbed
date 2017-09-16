package encoding

import (
	"fmt"
	"time"
)

func init() {
}

func timeToString(t time.Time) string {
	return fmt.Sprintf("%02d:%02d:%02d.%03d",
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond() / 1000)
}

// Crop takes file and start and length and sets to out file
func Crop(in string, out string, start time.Time, length time.Time) {
	Ffmpeg("-ss", timeToString(start),
		"-i", in,
		"-t", timeToString(length),
		"-c:v", "libx264", "-strict", "experimental", out)
}
