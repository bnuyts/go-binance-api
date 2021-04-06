package model

type Ticker struct {
	Symbol string  `json:"Symbol"`
	Price  float32 `json:"Price"`
}
