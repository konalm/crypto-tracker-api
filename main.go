package main

import (
	"crypto-tracker-api/bitcoinRates"
	"crypto-tracker-api/cryptoRatesController"
	"crypto-tracker-api/fetchCryptoRates"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"
	"log"
	"net/http"
)

func main() {
	cryptoRates := fetchCryptoRates.FetchCryptoRates()
	cryptoRatesController.InsertCryptoRates(cryptoRates)
	router := mux.NewRouter()
	router.HandleFunc("/rates", bitcoinRates.GetBitcoinRates).Methods("GET")
	c := cron.New()
	c.Start()
	c.AddFunc("0 */5 * * * *", func() {
		fmt.Println("Every 5th min")
		fmt.Println("call coin api every 5 mins")
		bitcoinRate := fetchCryptoRates.FetchBitcoinRate()
		bitcoinRates.InsertBitcoinRate(bitcoinRate)
	})

	originsAllowed := handlers.AllowedOrigins([]string{"http://localhost:8081"})
	headersAllowed := handlers.AllowedHeaders([]string{"X-Requested-With"})
	methodsAllowed := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":8484", handlers.CORS(originsAllowed, headersAllowed, methodsAllowed)(router)))
}
