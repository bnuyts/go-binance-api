package calculations

import (
	"log"

	"github.com/binance-exchange/go-binance"
)

func LastPrice(symbol string) float64 {
	apiService := binance.NewAPIService("https://www.binance.com", "", nil, nil, nil)
	client := binance.NewBinance(apiService)

	ticker, error := client.Ticker24(binance.TickerRequest{Symbol: symbol})

	if error != nil {
		log.Println(error)
		return 0
	}

	return ticker.LastPrice
}
