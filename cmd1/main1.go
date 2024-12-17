package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	log.Printf("server running on port ----------------------%s", port)

	r := gin.Default()

	r.GET("/what", func(c *gin.Context) {
		c.JSON(200, gin.H{"service": "22"})
	})

	err := r.Run(":" + port)

	if err != nil {
		panic(err)
	}
}
