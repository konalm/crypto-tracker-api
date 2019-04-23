package main

import (
  "log"
  "net/http"
  "github.com/gorilla/handlers"
  "stelita-api/router"
  "stelita-api/cronJobs"
  "stelita-api/env"
  "stelita-api/config"
  // "stelita-api/rsi"
  // "stelita-api/httpRequests"
  "stelita-api/errorReporter"

  // "stelita-api/cryptoRatesController"
  // "stelita-api/fetchCryptoRates"
)


func main() {
  env.SetEnvVariables()

  appRouter := router.Index()

  // httpRequests.StartAnalysisReports()
  // httpRequests.UpdateAnalysisReports()

  // cryptoRates := fetchCryptoRates.FetchCryptoRates()
  // cryptoRatesController.InsertCryptoRates(cryptoRates)
  // cryptoRates := fetchCryptoRates.FetchCryptoRatesFromCoinMarketCapApi()
  // cryptoRatesController.InsertCryptoRatesFromCoinMarketCapApi(cryptoRates)

  errorReporter.ReportError("Call Crons")
	cronJobs.HandleBitcoinRate()
  cronJobs.HandleRankedCryptoCurrencyUpdate()

  originsAllowed := handlers.AllowedOrigins([]string{config.ALLOWED_CLIENT, config.ALLOWED_CLIENT_2})
  headersAllowed := handlers.AllowedHeaders([]string{"pkm-client", "X-Requested-With", "content-type", "Authorization", "Access-Control-Allow-Origin"})
  methodsAllowed := handlers.AllowedMethods([]string{"POST", "HEAD", "GET", "PUT", "OPTIONS"})

  log.Fatal(http.ListenAndServe(":" + config.PORT, handlers.CORS(originsAllowed, headersAllowed, methodsAllowed)(appRouter)))
}
