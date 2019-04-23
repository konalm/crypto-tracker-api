package cronJobs

import (
  "github.com/robfig/cron"
  "stelita-api/cryptoRatesController"
  "stelita-api/fetchCryptoRates"
  "stelita-api/rsi"
  // "stelita-api/indicatorReporter"
  "stelita-api/httpRequests"
  "stelita-api/errorReporter"
)

/**
 *
 */
func HandleBitcoinRate() {
  errorReporter.ReportError("Handle Bitcoin Rate")

  c := cron.New()
  c.Start()

  c.AddFunc("0 */15 * * * *", func() {
    errorReporter.ReportError("Handle bitcoin rate on 15 min cron")

    // cryptoRates := fetchCryptoRates.FetchCryptoRates()
    // cryptoRatesController.InsertCryptoRates(cryptoRates)

    cryptoRates := fetchCryptoRates.FetchCryptoRatesFromCoinMarketCapApi()
    cryptoRatesController.InsertCryptoRatesFromCoinMarketCapApi(cryptoRates)

    rsi.HandleRsi()

    errorReporter.ReportError("Handled RSI")

    // indicatorReporter.ReportIndicatorsViaEmail(1)

    httpRequests.StartAnalysisReports()
    httpRequests.UpdateAnalysisReports()
  })
}
