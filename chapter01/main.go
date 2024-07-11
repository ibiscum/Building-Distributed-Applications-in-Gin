package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", IndexHandler)
	err := http.ListenAndServe(":8001", router)
	if err != nil {
		log.Fatal(err)
	}
}
