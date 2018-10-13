package fetchCryptoRates

import (
	"crypto-tracker-api/structs"
	. "encoding/json"
	. "fmt"
	. "io/ioutil"
	"net/http"
)

func FetchBitcoinRate() structs.BitcoinRate {
	client := &http.Client{}
	request, err := http.NewRequest(
		"GET",
		"https://rest.coinapi.io/v1/exchangerate/BTC/USD", nil,
	)
	request.Header.Set("X-CoinApi-Key", `E4C3D4AE-29D8-4A9F-BD36-EB367D836532`)

	resp, err := client.Do(request)
	if err != nil {
		Println("handle http error")
	}

	if resp.StatusCode != 200 {
		panic("ERROR fetching bitcoin rate from coinapi")
	}

	body, err := ReadAll(resp.Body)
	if err != nil {
		Println("ERROR reading response body")
	}
	jsonBody := string(body)
	var bitcoinRate structs.BitcoinRate
	Unmarshal([]byte(jsonBody), &bitcoinRate)

	return bitcoinRate
}
