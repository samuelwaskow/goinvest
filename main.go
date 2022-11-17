package main

import (
	"net/http"
	"strconv"

	"com.stocks/investing/model"
	"github.com/gin-gonic/gin"
)

var stocks = []model.Stock{
	{ID: 1, Name: "Energias do Brasil", Ticker: "ENBR3"},
	{ID: 2, Name: "Taurus Armas", Ticker: "TASA3"},
	{ID: 3, Name: "Taurus Armas", Ticker: "TASA3"},
}

func main() {
	router := gin.Default()

	router.GET("/stock/:id", getStock)
	router.GET("/stock", listStock)
	router.POST("/stock", createStock)
	router.PUT("/stock/:id", updateStock)
	router.DELETE("/stock/:id", deleteStock)

	router.Run(":8080")
}

func getStock(c *gin.Context) {
	c.JSON(http.StatusOK, stocks)
}

func listStock(c *gin.Context) {
	c.JSON(http.StatusOK, stocks)
}

func createStock(c *gin.Context) {
	var newStock model.Stock

	if err := c.BindJSON(&newStock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	stocks = append(stocks, newStock)
	println("stocks - supr", len(stocks))
	c.JSON(http.StatusCreated, newStock)
}

func updateStock(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var stock model.Stock

	if err := c.BindJSON(&stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	index := -1
	for i := 0; i < len(stocks); i++ {
		if stocks[i].ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Stock not found",
		})
		return
	}
	stocks[index] = stock

	c.JSON(http.StatusOK, stocks)
}

func deleteStock(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	index := -1
	for i := 0; i < len(stocks); i++ {
		if stocks[i].ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Stock not found",
		})
		return
	}
	stocks = append(stocks[:index], stocks[index+1:]...)
	c.JSON(http.StatusOK, gin.H{
		"message": "Company has been deleted",
	})
}
