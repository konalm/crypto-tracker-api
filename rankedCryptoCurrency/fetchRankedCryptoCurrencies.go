package rankedCryptoCurrency

import (
  "net/http"
  "fmt"
  "io/ioutil"
  "encoding/json"
  "stelita-api/reports"
  "stelita-api/errorReporter"
)

type CoinMarketGapApiResponse struct {
  Data map[string] RankedCryptoCurrency
}


/**
 *
 */
func FetchRankedCryptoCurrencies() map[string] RankedCryptoCurrency {
  client := &http.Client{}
  url := "https://api.coinmarketcap.com/v2/ticker/?limit=100"
  request, err := http.NewRequest("GET", url, nil)

  resp, err := client.Do(request)
  if err != nil {
    fmt.Println(err.Error())
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    panic("Error reading response body")
  }

  var responseMessage string
  if resp.StatusCode != 200 {
    errorReporter.ReportError("Fetching ranked crypto currencies from Coin API")
    responseMessage = string(body)
  }

  jsonBody := string(body)
  var apiResponse CoinMarketGapApiResponse
  json.Unmarshal([]byte(jsonBody), &apiResponse)

  reports.InsertFetchRankedCryptoCurrenciesReports(resp.StatusCode, responseMessage)

  return apiResponse.Data
}
