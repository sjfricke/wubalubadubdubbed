package encoding

import (
	"bytes"
	"strconv"
)

func init() {
}

// Crop takes file and start and length and sets to out file
func Stitch(inputs []string, output string) error {
	var stitchArgs []string

	// add input files
	for _, in := range inputs {
		stitchArgs = append(stitchArgs, "-i", in)
	}

	stitchArgs = append(stitchArgs, "-filter_complex")

	var filBuf bytes.Buffer
//	filBuf.WriteString("'")
	for i, _ := range inputs {
		filBuf.WriteString("[")
		filBuf.WriteString(strconv.Itoa(i))
		filBuf.WriteString(":0][")
		filBuf.WriteString(strconv.Itoa(i))
		filBuf.WriteString(":1]")
	}

	filBuf.WriteString(" concat=n=")
	filBuf.WriteString(strconv.Itoa(len(inputs)))
	filBuf.WriteString(":v=1:a=1 [v][a]")
	stitchArgs = append(stitchArgs, filBuf.String())

	stitchArgs = append(stitchArgs, "-map", "[v]", "-map", "[a]")
	stitchArgs = append(stitchArgs, "-strict", "-2", output);
	
	err := Ffmpeg(stitchArgs...)
	return err
}
