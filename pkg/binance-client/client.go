package binanceclient

import (
	"sync"

	"github.com/binance-exchange/go-binance"
)

var (
	createClient sync.Once
	client       binance.Binance
)

func GetClient() binance.Binance {
	createClient.Do(func() {
		apiService := binance.NewAPIService("https://www.binance.com", "", nil, nil, nil)
		client = binance.NewBinance(apiService)
	})

	return client
}
