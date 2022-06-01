package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

type Products struct {
	Produtos []Product `json:"produtos"`
}

type Product struct {
	Id         string  `json:"id"`
	Nome       string  `json:"nome"`
	Cor        string  `json:"cor"`
	Preco      float64 `json:"preco"`
	Estoque    int     `json:"estoque"`
	Codigo     string  `json:"codigo"`
	Publicacao string  `json:"publicacao"`
	Criacao    int64   `json:"criacao"`
}

func handler(c *gin.Context) {

	productId := c.Param("id")

	jsonFile, err := os.Open("products.json")

	if err != nil {
		c.JSON(500, gin.H{
			"message": "error on open file",
		})
		fmt.Println(err)
		return
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var products Products

	json.Unmarshal(byteValue, &products)

	for i := 0; i < len(products.Produtos); i++ {
		if products.Produtos[i].Id == productId {
			c.JSON(200, products.Produtos[i])
			return
		}
	}

	c.JSON(404, gin.H{
		"message": "Product not found",
	})

	defer jsonFile.Close()

}

func main() {
	router := gin.Default()

	router.GET("/products/:id", handler)

	router.Run()
}
