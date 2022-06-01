package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func handler(c *gin.Context) {

	fmt.Println(c.Param("id"))

	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func main() {
	router := gin.Default()

	router.GET("/products/:id", handler)

	router.Run()
}
