package main

import (
	"log"
	"net/http"

	"github.com/bnuyts/go-binance-api/pkg/api"
	"github.com/gorilla/mux"
)

func setup() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/mcap/{symbol}", api.LastPrice)

	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	log.Println("Starting API")
	setup()
}
