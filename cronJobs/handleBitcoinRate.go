package cronJobs

import (
  "github.com/robfig/cron"
  "stelita-api/cryptoRatesController"
  "stelita-api/fetchCryptoRates"
  "stelita-api/rsi"
  "stelita-api/indicatorReporter"
)

/**
 *
 */
func HandleBitcoinRate() {
  c := cron.New()
  c.Start()

  c.AddFunc("0 */15 * * * *", func() {
    cryptoRates := fetchCryptoRates.FetchCryptoRates()
    cryptoRatesController.InsertCryptoRates(cryptoRates)
    rsi.HandleRsi()
    indicatorReporter.ReportIndicatorsViaEmail(1)
  })
}
