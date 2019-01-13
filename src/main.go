package main

import (
	"net/http"

	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()

	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong") //string
	})

	engine.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome Ayman!"}) //JSON with Helper map function
	})
	//run server on PORT
	engine.Run(port())
}

func port() string {

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8080"
	}

	return ":" + port
}
