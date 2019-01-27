package fetchCryptoRates

import (
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
