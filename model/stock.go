package model

type Stock struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
}
