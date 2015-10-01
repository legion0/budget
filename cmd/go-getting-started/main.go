package main

import (
	_ "fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/oauth-io/sdk-go"

	"github.com/heroku/go-getting-started/cmd/go-getting-started/dal"
)

func main() {
	//var err error
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	dal.Connect()

	router := gin.New()
	router.Use(gin.Logger())
	router.Static("/static", "static")
	router.Static("/components", "components")

	router.StaticFile("/", "./components/main/main.html")
	//router.StaticFile("/components/main/controller.js", "./components/main/controller.js")
	//router.StaticFile("/components/main/main.css", "./components/main/main.css")

	oauth := oauthio.New(os.Getenv("OAUTHIO_PUBLIC_KEY"), os.Getenv("OAUTHIO_SECRET_KEY"))

	router.GET("/signin", func(c *gin.Context) {
		oauth.Redirect("google", "http://localhost:5000/")
	})

	router.Run(":" + port)
}
