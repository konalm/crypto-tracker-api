package router

import (
  "flag"
  "net/http"
  "github.com/gorilla/mux"
  "stelita-api/bitcoinRates"
  "stelita-api/cryptoRatesController"
  "stelita-api/rankedCryptoCurrency"
  "stelita-api/authentication"
  "stelita-api/middleware"
  "stelita-api/reports"
)


func Index() *mux.Router {
  router := mux.NewRouter()

  router.HandleFunc("/bitcoin-rates", bitcoinRates.GetBitcoinRates).Methods("GET")
  router.HandleFunc("/crypto-currencies", cryptoRatesController.GetCryptoCurrencies).Methods("GET")
  router.HandleFunc("/crypto-rates", cryptoRatesController.GetCryptoCurrencyRates).Methods("GET")
  router.HandleFunc("/crypto-data", rankedCryptoCurrency.GetCryptoCurrencyData).Methods("GET")

  router.HandleFunc("/login", authentication.Login).Methods("POST")
  router.HandleFunc("/auth-verification", authentication.AuthCheck).Methods("GET")

  authRouter := router.NewRoute().Subrouter()
  authRouter.HandleFunc("/protected", authentication.ProtectedResource).Methods("GET")
  authRouter.HandleFunc("/crypto-data-reports", reports.GetCryptoDataReports).Methods("GET")
  authRouter.HandleFunc("/ranked-crypto-currencies-reports", reports.GetRankedCryptoCurrenciesReports).Methods("GET")
  authRouter.HandleFunc("/insert-crypto-reports", reports.GetInsertedCryptoReports).Methods("GET")
  authRouter.HandleFunc("/update-crypto-trend-stat-reports", reports.GetUpdateCryptoTrendStatReports).Methods("GET")
  authRouter.Use(middleware.Auth)

  var dir string
  flag.StringVar(&dir, "dir", ".", "assets")
  flag.Parse()

  /* serve static files */
  router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

  return router
}
