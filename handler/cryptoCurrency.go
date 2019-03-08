package handler


import (
  "fmt"
  "net/http"
  "github.com/gorilla/context"
  "encoding/json"
  // "strings"
  _ "github.com/go-sql-driver/mysql"
  "stelita-api/rankedCryptoCurrency"
  "stelita-api/walletCurrency"
  "stelita-api/wallet"
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
