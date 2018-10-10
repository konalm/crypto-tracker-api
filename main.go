package main

import (
  "fmt"
  "net/http"
  "log"
  "github.com/robfig/cron"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
  "rest/fetchCryptoRates"
  "rest/bitcoinRates"
  "rest/cryptoRatesController"
)


/**
 *
 */
func main() {

  cryptoRates := fetchCryptoRates.FetchCryptoRates()
  cryptoRatesController.InsertCryptoRates(cryptoRates)

  // fmt.Println(cryptoRates)



  router := mux.NewRouter()
  router.HandleFunc("/rates", bitcoinRates.GetBitcoinRates).Methods("GET")

  c := cron.New()
  c.Start()

  c.AddFunc("0 */5 * * * *", func() {
    fmt.Println("Every 5th min")
    fmt.Println("call coin api every 5 mins")

    bitcoinRate := fetchCryptoRates.FetchBitcoinRate()
    bitcoinRates.InsertBitcoinRate(bitcoinRate)

    // cryptoRates := fetchCryptoRates.FetchCryptoRates()
    //
    // fmt.Println(cryptoRates)
  })

  originsAllowed := handlers.AllowedOrigins([]string{"http://localhost:8081"})
  headersAllowed := handlers.AllowedHeaders([]string{"X-Requested-With"})
  methodsAllowed := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

  log.Fatal(http.ListenAndServe(":8484", handlers.CORS(originsAllowed, headersAllowed, methodsAllowed)(router)))
}
