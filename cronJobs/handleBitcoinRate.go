package cronJobs

import (
  "fmt"
  "github.com/robfig/cron"
  "stelita-api/cryptoRatesController"
  "stelita-api/fetchCryptoRates"
  "stelita-api/rsi"
)

func HandleBitcoinRate() {
  fmt.Println("HANDLE BITCOIN RATE FUNC")

  c := cron.New()
  c.Start()

  c.AddFunc("* */15 * * * *", func() {
    fmt.Println("Run Cron")

    cryptoRates := fetchCryptoRates.FetchCryptoRates()
    cryptoRatesController.InsertCryptoRates(cryptoRates)

    rsi.HandleRsi()
  })
}
