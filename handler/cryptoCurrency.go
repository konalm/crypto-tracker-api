package handler


import (
  "fmt"
  "net/http"
  "github.com/gorilla/context"
  "encoding/json"
  // "strings"
  _ "github.com/go-sql-driver/mysql"
  "github.com/gorilla/mux"
  "stelita-api/rankedCryptoCurrency"
  "stelita-api/walletCurrency"
  "stelita-api/wallet"
  "stelita-api/cryptoCurrency"
  "stelita-api/cryptoRatesController"
  "stelita-api/abstractRatesByTimePeriod"
  "stelita-api/rsi"
)


/**
 *
 */
func GetCryptoCurrencyData(w http.ResponseWriter, r *http.Request) {
  fmt.Println("get crypto currency data !!")

  userIdInterface := context.Get(r, "userId")
  var userId int = userIdInterface.(int)

  wallet :=  wallet.GetUsersWallet(userId)
  currenciesInWallet := walletCurrency.GetWalletCurrenciesForWallet(wallet.Id)
  cryptoCurrencyData := rankedCryptoCurrency.GetCryptoCurrencyData(currenciesInWallet)

  json.NewEncoder(w).Encode(cryptoCurrencyData)
}

/**
 *
 */
func GetCryptoCurrencyItemData(w http.ResponseWriter, r *http.Request) {
  fmt.Println("get crypto currency item data")
  params := mux.Vars(r)
  symbol := params["symbol"]

  cryptoCurrencyData := cryptoCurrency.GetCryptoCurrencyData(symbol)
  if cryptoCurrencyData.Symbol == "" {
    w.WriteHeader(404)
    w.Write([]byte("Crypto currency not found"))
    return
  }

  fmt.Println("crypto currency is found")

  rates := cryptoRatesController.GetCryptoCurrencyRatesForRsi(symbol)
  fmt.Println("got rates")
  ratesIn15MinPeriod := abstractRatesByTimePeriod.FifteenMinPeriods(rates)
  fmt.Println("got rates in 15min period")

  cryptoCurrencyData.RSI = rsi.CalculateRsi(ratesIn15MinPeriod, 0)
  cryptoCurrencyData.Rates = ratesIn15MinPeriod [0:15]
  cryptoCurrencyData.ClosingPrice = cryptoCurrencyData.Rates[0].ClosingPrice

  json.NewEncoder(w).Encode(cryptoCurrencyData)
}
