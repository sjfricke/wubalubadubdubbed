package main

import (
	"fmt"
	"github.com/sjfricke/wubalubadubdub/database"
	"github.com/sjfricke/wubalubadubdub/encoding"
	"time"
)

func main() {
	fmt.Println("hello")

	database.Test();

	start := time.Date(2009, time.November, 10, 0, 0, 2, 418000, time.UTC)
	end := time.Date(2009, time.November, 10, 0, 0, 0, 453000, time.UTC)

	encoding.Crop("/home/fricke/Videos/rm1.mp4", "/home/fricke/Videos/test.mp4", start, end);
}
