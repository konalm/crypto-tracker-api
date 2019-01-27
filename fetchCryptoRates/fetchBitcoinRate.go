package fetchCryptoRates

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
  "stelita-api/structs"
  "stelita-api/errorReporter"
)


/**
 *
 */
func FetchBitcoinRate() structs.BitcoinRate {
	client := &http.Client{}
	request, err := http.NewRequest(
		"GET",
		"https://rest.coinapi.io/v1/exchangerate/BTC/USD", nil,
	)
	request.Header.Set("X-CoinApi-Key", `E4C3D4AE-29D8-4A9F-BD36-EB367D836532`)

	resp, err := client.Do(request)
	if err != nil {
    errorReporter.ReportError("Http request to Coin API")
	}

	if resp.StatusCode != 200 {
    errorReporter.ReportError("Fetching bitcoin rate from Coin API")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
    errorReporter.ReportError("Reading data from fetching bitcoin rate from  Coin API")
	}
	jsonBody := string(body)
	var bitcoinRate structs.BitcoinRate
	json.Unmarshal([]byte(jsonBody), &bitcoinRate)

	return bitcoinRate
}
