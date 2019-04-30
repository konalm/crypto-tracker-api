package router

import (
  "flag"
  "net/http"
  "github.com/gorilla/mux"
  "stelita-api/handler"
  "stelita-api/bitcoinRates"
  "stelita-api/cryptoRatesController"
  "stelita-api/authentication"
  "stelita-api/middleware"
)


func Index() *mux.Router {
  router := mux.NewRouter()

  router.HandleFunc("/bitcoin-rates", bitcoinRates.GetBitcoinRates).Methods("GET")
  router.HandleFunc("/crypto-currencies", cryptoRatesController.GetCryptoCurrencies).Methods("GET")
  router.HandleFunc("/crypto-rates", cryptoRatesController.GetCryptoCurrencyRates).Methods("GET")
  router.HandleFunc("/login", authentication.Login).Methods("POST", "OPTIONS")
  router.HandleFunc("/auth-verification", authentication.AuthCheck).Methods("GET")

  router.HandleFunc("/analysis", handler.GetAnalysis).Methods("GET")
  router.HandleFunc("/crypto-currency/{crypto_symbol}/analysis", handler.GetCryptoCurrencyAnalysis).Methods("GET")
  router.HandleFunc("/analysis/{id}", handler.GetAnalysisItem).Methods("GET")

  authRouter := router.NewRoute().Subrouter()
  authRouter.HandleFunc("/crypto-data", handler.GetCryptoCurrencyData).Methods("GET")
  authRouter.HandleFunc("/setup-wallet", handler.SetupWallet).Methods("POST")
  authRouter.HandleFunc("/transaction", handler.CreateTransaction).Methods("POST")
  authRouter.HandleFunc("/protected", authentication.ProtectedResource).Methods("GET")
  authRouter.HandleFunc("/wallet-states", handler.GetUserWalletStates).Methods("GET")
  authRouter.HandleFunc("/user-transactions", handler.GetUserTransactions).Methods("GET")
  authRouter.HandleFunc("/currencies-in-wallet", handler.GetCurrenciesInWallet).Methods("GET")
  authRouter.HandleFunc("/user-wallet-currencies", handler.GetWalletCurrenciesForWallet).Methods("GET")
  authRouter.HandleFunc("/users-latest-wallet", handler.GetUserLatestWalletState).Methods("GET")
  authRouter.HandleFunc("/event-reports/{event_id}", handler.GetEventReports).Methods("GET")
  authRouter.HandleFunc("/crypto-currency/{symbol}", handler.GetCryptoCurrencyItemData).Methods("GET")
  authRouter.HandleFunc("/crypto-currency/{symbol}/closing_prices/{data_count}", handler.GetCryptoRates).Methods("GET")
  authRouter.Use(middleware.Auth)

  var dir string
  flag.StringVar(&dir, "dir", ".", "assets")
  flag.Parse()

  /* serve static files */
  router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

  return router
}
