package router

import (
  "github.com/gorilla/mux"
  "stelita-api/bitcoinRates"
  "stelita-api/cryptoRatesController"
  "stelita-api/rankedCryptoCurrency"
  "net/http"
  "flag"
)


func Index() *mux.Router {
  router := mux.NewRouter()
  router.HandleFunc("/bitcoin-rates", bitcoinRates.GetBitcoinRates).Methods("GET")
  router.HandleFunc("/crypto-currencies", cryptoRatesController.GetCryptoCurrencies).Methods("GET")
  router.HandleFunc("/crypto-rates", cryptoRatesController.GetCryptoCurrencyRates).Methods("GET")
  router.HandleFunc("/crypto-data", rankedCryptoCurrency.GetCryptoCurrencyData).Methods("GET")

  var dir string
  flag.StringVar(&dir, "dir", ".", "assets")
  flag.Parse()

  /* serve static files */
  router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

  return router
}
