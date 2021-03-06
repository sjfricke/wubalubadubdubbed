package main

import (
	"github.com/sjfricke/wubalubadubdub/database"
	"github.com/sjfricke/wubalubadubdub/encoding"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/fatih/set"
	"fmt"
	"strings"
	"os/exec"
	"path/filepath"
	"log"
)

type PostData struct {
	Data     string `json:"data" binding:"required"`
}

type LastPost struct {
	URL string
	Phrase string
}

// THIS CODE IS STILL BROKEN
// Currently all databse file paths are toLower() and
// causes file paths wit uppercase to fail ffmpeg part
func main() {
	db := database.ConnectCockroach("postgresql://root@localhost:26257?sslmode=disable")

	last := LastPost{ "/public/123.mp4", "wubalubadubdub" }
	
	router := gin.Default()

	router.LoadHTMLFiles("index.tmpl")
	
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", last)
	})

	router.POST("/", func(c *gin.Context) {
		var json PostData
		var extra int = 0

		if c.BindJSON(&json) != nil {
			c.JSON(http.StatusBadRequest, gin.H{"BAD DATA": "the post body so please add it :)"})
		}
		if json.Data == "" {
			c.JSON(http.StatusBadRequest, gin.H{"BAD DATA": "The data key of the post body"})
		}
		text := strings.TrimSpace(string(json.Data))
		log.Println("Phrase: ", text)
		words := strings.Split(strings.ToLower(text), " ")
		phraseEntries := make([]database.PhraseEntry, len(words))
		missing := set.New()
		for i := 0; i < len(words); i++ {
			phraseEntries[i], extra = database.ReadPhrase(db, words[i:]... )
			if phraseEntries[i].Phrase == "" {
				// phrase not found
				missing.Add(words[i])
			}

			if extra > 1 {
				i = i + (extra-1)
			}
		}
		if(missing.IsEmpty()) {
			dir, err := encoding.Encode(phraseEntries)
			if err != nil {
				c.String(http.StatusInternalServerError, "Sorry, Morty messed up and Rick needs to fix this, please try again")
			}
			
			orig := filepath.Join(dir, "output.mp4")
			out := filepath.Join("public", strings.Join([]string{dir, "mp4"}, "."))
			log.Println("New video at: ", out)
			exec.Command("mv", orig, out).Run()
			exec.Command("rm", "-r", dir).Run()
			last.URL = fmt.Sprintf("http://wubalubadubdubbed.com/%s", out)
			last.Phrase = text;
			c.JSON(http.StatusOK, gin.H{"url": last.URL})
		} else {
			log.Println("Failed words: ", missing)
			c.JSON(http.StatusBadRequest, gin.H{"missing": missing.List()})
		}
	})

	router.Static("/public", "public")

	router.Run(":8000") // listen and serve on 0.0.0.0:8080
}
