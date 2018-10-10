package fetchCryptoRates

import (
  "net/http"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "rest/structs"
)


type ApiResponse struct {
  Asset_id_base string
  Rates []structs.BitcoinRate
}


/**
 *
 */
func FetchCryptoRates() []structs.BitcoinRate {
  fmt.Println("Fetch crypto rates from coinapi")

  client := &http.Client{}

  request, err :=
    http.NewRequest("GET", "https://rest.coinapi.io/v1/exchangerate/BTC", nil)
  request.Header.Set("X-CoinApi-Key", `E4C3D4AE-29D8-4A9F-BD36-EB367D836532`)

  resp, err := client.Do(request)
  if err != nil {
    fmt.Println(err.Error())
  }

  if resp.StatusCode != 200 {
    panic("ERROR fetching crypto rates")
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Println("ERROR reading response body")
  }

  jsonBody := string(body)
  var apiResponse ApiResponse

  json.Unmarshal([]byte(jsonBody), &apiResponse)

  //
  // var cryptoRates []structs.BitcoinRate
  //
  // for _, cryptoRate := range apiResponse.Rates {
  //   fmt.Println("in foreach rate of api response >>>> ")
  //   fmt.Println(cryptoRate)
  // }

  return apiResponse.Rates
}
