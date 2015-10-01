package main

import (
	_ "fmt"
	_ "github.com/heroku/go-getting-started/cmd/go-getting-started/dal"
	"log"
	"os"

	"github.com/heroku/go-getting-started/Godeps/_workspace/src/github.com/gin-gonic/gin"
	_ "github.com/heroku/go-getting-started/Godeps/_workspace/src/github.com/lib/pq"
)

func main() {
	//var err error
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Static("/static", "static")
	router.Static("/components", "components")

	router.StaticFile("/", "./components/main/main.html")

	router.Run(":" + port)
}
