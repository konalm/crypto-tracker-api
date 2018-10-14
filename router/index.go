package router

import (
  "github.com/gorilla/mux"
  "crypto-tracker-api/bitcoinRates"
  "crypto-tracker-api/cryptoRatesController"
)


func Index() *mux.Router {
  router := mux.NewRouter()
  router.HandleFunc("/bitcoin-rates", bitcoinRates.GetBitcoinRates).Methods("GET")
  router.HandleFunc("/crypto-currencies", cryptoRatesController.GetCryptoCurrencies).Methods("GET")
  router.HandleFunc("/crypto-rates", cryptoRatesController.GetCryptoCurrencyRates).Methods("GET")

  return router
}
