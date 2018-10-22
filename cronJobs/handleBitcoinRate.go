package cronJobs

import (
  "fmt"
  "github.com/robfig/cron"
  "crypto-tracker-api/cryptoRatesController"
  "crypto-tracker-api/fetchCryptoRates"
  "crypto-tracker-api/rsi"
)

func HandleBitcoinRate() {
  c := cron.New()
  c.Start()

  c.AddFunc("0 */15 * * * *", func() {
    fmt.Println("Run Cron")

    cryptoRates := fetchCryptoRates.FetchCryptoRates()
    cryptoRatesController.InsertCryptoRates(cryptoRates)

    rsi.HandleRsi()
  })
}
