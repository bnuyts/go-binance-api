package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bnuyts/go-binance-api/pkg/api"
	"github.com/gorilla/mux"
)

func help(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/static/help.html")
}

func setup() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "10000"
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", help)
	router.HandleFunc("/price/last/{symbol}", api.LastPrice)

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func main() {
	log.Println("Starting API")
	setup()
}
