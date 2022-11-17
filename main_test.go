package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"com.stocks/investing/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const ENDPOINT = "/stock"

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestNewStockHandler(t *testing.T) {
	r := SetUpRouter()
	r.POST(ENDPOINT, createStock)

	stock := model.Stock{
		ID:     10,
		Name:   "Demo Company",
		Ticker: "Demo CEO",
	}
	jsonValue, _ := json.Marshal(stock)
	req, _ := http.NewRequest("POST", ENDPOINT, bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestListStockHandler(t *testing.T) {
	r := SetUpRouter()
	r.GET(ENDPOINT, listStock)
	req, _ := http.NewRequest("GET", ENDPOINT, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var stocks []model.Stock
	json.Unmarshal(w.Body.Bytes(), &stocks)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, stocks)
}

func TestUpdateStockHandler(t *testing.T) {
	r := SetUpRouter()
	r.PUT(ENDPOINT+"/:id", updateStock)
	stock := model.Stock{
		ID:     2,
		Name:   "Taesa",
		Ticker: "TAEE3",
	}
	jsonValue, _ := json.Marshal(stock)
	reqFound, _ := http.NewRequest("PUT", ENDPOINT+"/"+fmt.Sprint(stock.ID), bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusOK, w.Code)

	reqNotFound, _ := http.NewRequest("PUT", ENDPOINT+"/12", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqNotFound)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteStockHandler(t *testing.T) {
	r := SetUpRouter()
	r.DELETE(ENDPOINT+"/:id", deleteStock)

	reqFound, _ := http.NewRequest("DELETE", ENDPOINT+"/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusOK, w.Code)

	reqNotFound, _ := http.NewRequest("DELETE", ENDPOINT+"/12", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqNotFound)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
