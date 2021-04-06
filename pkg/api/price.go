package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bnuyts/go-binance-api/pkg/calculations"
	"github.com/gorilla/mux"
)

func LastPrice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	symbol := strings.ToUpper(vars["symbol"])

	fmt.Fprint(w, calculations.LastPrice(symbol))
}
