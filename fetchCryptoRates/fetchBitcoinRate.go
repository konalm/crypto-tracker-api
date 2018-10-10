package fetchCryptoRates 

import (
  "net/http"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "rest/structs"
)


/**
 *
 */
func FetchBitcoinRate() structs.BitcoinRate {
  client := &http.Client{}

  request, err :=
    http.NewRequest("GET", "https://rest.coinapi.io/v1/exchangerate/BTC/USD", nil)
  request.Header.Set("X-CoinApi-Key", `E4C3D4AE-29D8-4A9F-BD36-EB367D836532`)

  resp, err := client.Do(request)
  if err != nil {
    fmt.Println("handle http error")
  }

  if resp.StatusCode != 200 {
    panic("ERROR fetching bitcoin rate from coinapi")
    // return nil
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Println("ERROR reading response body")
  }

  jsonBody := string(body)
  var bitcoinRate structs.BitcoinRate

  fmt.Println("fetch bitcoin rate response >>")
  fmt.Println(jsonBody)

  json.Unmarshal([]byte(jsonBody), &bitcoinRate)

  return bitcoinRate
}
