package encoding

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func init() {
	cmd := exec.Command("ffmpeg", "-version")

	err := cmd.Run()
	if err != nil {
		fmt.Printf("ffmpeg not installed or set in path\n")
		os.Exit(-1)
	}
}

// Ffmpeg is used to execute ffmpeg in shell
func Ffmpeg(args ...string) error{
	cmd := exec.Command("ffmpeg", args...)

	fmt.Printf("ffmpeg %s\n", strings.Join(args, " "))

	err := cmd.Run()

	if err != nil {
		fmt.Printf("Failed ffmpeg %s: %v\n", strings.Join(args, " "), err)
	}

	return err
}
