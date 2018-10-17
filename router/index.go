package router

import (
  "github.com/gorilla/mux"
  "crypto-tracker-api/bitcoinRates"
  "crypto-tracker-api/cryptoRatesController"
  "net/http"
  "flag"
)


func Index() *mux.Router {
  router := mux.NewRouter()
  router.HandleFunc("/bitcoin-rates", bitcoinRates.GetBitcoinRates).Methods("GET")
  router.HandleFunc("/crypto-currencies", cryptoRatesController.GetCryptoCurrencies).Methods("GET")
  router.HandleFunc("/crypto-rates", cryptoRatesController.GetCryptoCurrencyRates).Methods("GET")

  var dir string
  flag.StringVar(&dir, "dir", ".", "assets")
  flag.Parse()

  /* serve static files */
  router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
  return router
}
