package main

import (
	"github.com/sjfricke/wubalubadubdub/database"
//	"github.com/sjfricke/wubalubadubdub/encoding"
	"github.com/gin-gonic/gin"
	"net/http"
	// "time"
	"io/ioutil"
	"github.com/fatih/set"
	// "fmt"
	"strings"
)

func main() {
	db := database.ConnectCockroach("postgresql://root@localhost:26257?sslmode=disable")

	router := gin.Default()

	router.POST("/", func(c *gin.Context) {
		bytes, _ := ioutil.ReadAll(c.Request.Body)
		text := string(bytes)
		words := strings.Split(strings.ToLower(text), " ")
		phraseEntries := make([]*database.PhraseEntry, len(words))
		missing := set.New()
		for i, w := range words {
			phraseEntries[i] = database.ReadPhrase(db, w)
			if(phraseEntries[i].Phrase == "") {
				// phrase not found
				missing.Add(w)
			}
		}
		if(missing.IsEmpty()) {
			c.JSON(http.StatusOK, gin.H{"text": "beep boop"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"missing": missing.List()})
		}
	})

	router.Run(":8000") // listen and serve on 0.0.0.0:8080
}
