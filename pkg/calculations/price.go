package calculations

import (
	"log"

	"github.com/binance-exchange/go-binance"
	binanceclient "github.com/bnuyts/go-binance-api/pkg/binance-client"
)

func LastPrice(symbol string) float64 {
	ticker, error := binanceclient.GetClient().Ticker24(binance.TickerRequest{Symbol: symbol})

	if error != nil {
		log.Println(error)
		return 0
	}

	return ticker.LastPrice
}
