package main

import (
	"github.com/sjfricke/wubalubadubdub/database"
	"github.com/sjfricke/wubalubadubdub/encoding"
	"github.com/gin-gonic/gin"
	"net/http"
	// "time"
	"github.com/fatih/set"
	"fmt"
	"strings"
	"os/exec"
	"path/filepath"
)

type PostData struct {
	Data     string `json:"data" binding:"required"`
}

// THIS CODE IS STILL BROKEN
// Currently all databse file paths are toLower() and
// causes file paths wit uppercase to fail ffmpeg part
func main() {
	db := database.ConnectCockroach("postgresql://root@localhost:26257?sslmode=disable")

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"TEST":"Delete This"})
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
		text := string(json.Data)
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
			dir := encoding.Encode(phraseEntries)
			fmt.Println(dir)
			orig := filepath.Join(dir, "output.mp4")
			fmt.Println(orig)
			out := filepath.Join("public", strings.Join([]string{dir, "mp4"}, "."))
			fmt.Println(out)
			exec.Command("mv", orig, out).Run()
			exec.Command("rm", "-r", dir).Run()
			c.JSON(http.StatusOK, gin.H{"url": fmt.Sprintf("http://wubalubadubdubbed.com/%s", out)})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"missing": missing.List()})
		}
	})

	router.Static("/public", "public")

	router.Run(":8000") // listen and serve on 0.0.0.0:8080
}
