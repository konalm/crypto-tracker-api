package fetchCryptoRates

import (
  "fmt"
  "net/http"
  "encoding/json"
  "io/ioutil"
  "stelita-api/structs"
  "stelita-api/reports"
  "stelita-api/errorReporter"
)


type ApiResponse struct {
  Asset_id_base string
  Rates []structs.BitcoinRate
}

/**
 *
 */
func FetchCryptoRates() []structs.BitcoinRate {
  client := &http.Client{}
  request, err := http.NewRequest("GET", "https://rest.coinapi.io/v1/exchangerate/BTC", nil)
  request.Header.Set("X-CoinApi-Key", `E4C3D4AE-29D8-4A9F-BD36-EB367D836532`)

  resp, err := client.Do(request)
  if err != nil {
    errorReporter.ReportError("Http request to get crypto rates from Coin API")
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    errorReporter.ReportError("Reading data from Coin API for crypto rates")
  }

  jsonBody := string(body)
  var apiResponse ApiResponse
  json.Unmarshal([]byte(jsonBody), &apiResponse)

  var responseMessage string
  if (resp.StatusCode != 200) {
    errorReporter.ReportError("Getting data from Coin API for crypto rates")
    responseMessage = string(body)
  }

  reports.InsertFetchCryptoDataReport("coinapi", resp.StatusCode, responseMessage)

  return apiResponse.Rates
}



/**
 *
 */
func FetchCryptoRatesFromCoinMarketCapApi() []CoinMarketCapCryptoCurrency {
  fmt.Println("fetch crypto rates from coin market cap api !!")

  client := &http.Client{}
  request, err := http.NewRequest(
    "GET",
    "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?convert=BTC",
    nil,
  )
  request.Header.Set("X-CMC_PRO_API_KEY", `a65f7aff-3f3b-4471-8d62-63fd0b7729a6`)

  response, err := client.Do(request)
  if err != nil {
    errorReporter.ReportError("Http request to get crypto rates from Coin Market Cap")
  }

  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    errorReporter.ReportError("Reading data from Coin Market Cap API for crypto rates")
  }

  jsonBody := string(body)
  var apiResponse CoinMarketCapApiResponse
  json.Unmarshal([]byte(jsonBody), &apiResponse)

  var responseMessage string
  if (response.StatusCode != 200) {
    errorReporter.ReportError("Error Getting data from Coin Market Cap API")
    responseMessage = string(body)
  }

  reports.InsertFetchCryptoDataReport("coinmarketcapapi", response.StatusCode, responseMessage)

  return apiResponse.Data
}
