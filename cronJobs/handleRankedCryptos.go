package cronJobs

import (
  "fmt"
  "github.com/robfig/cron"
  // "crypto-tracker-api/fetchCryptoRates"
  "crypto-tracker-api/rankedCryptoCurrency"
  "crypto-tracker-api/structs"
)


/**
 * Updated ranked crypto currencies everyday
 */
func updateRankedCryptos(cryptoCurrencies structs.) {
  c := cron.New()
  c.start()

  c.AddFunc("@daily", func() {
    rankedCryptos := rankedCryptoCurrency.FetchRankedCryptoCurrencies()
    rankedCryptoCurrency.InsertRankedCryptos(rankedCryptos)
  })
}
