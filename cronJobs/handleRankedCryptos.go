package cronJobs

import (
  "github.com/robfig/cron"
  "stelita-api/rankedCryptoCurrency"
)


/**
 * Updated ranked crypto currencies everyday
 */
func HandleRankedCryptoCurrencyUpdate() {
  c := cron.New()
  c.Start()

  c.AddFunc("@daily", func() {
    rankedCryptoCurrency.DestroyCurrentRankedCryptoCurrencies()
    rankedCryptos := rankedCryptoCurrency.FetchRankedCryptoCurrencies()
    rankedCryptoCurrency.InsertRankedCryptoCurrencies(rankedCryptos)
  })
}
