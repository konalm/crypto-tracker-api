package handler

import (
  "net/http"
  "encoding/json"
  "github.com/gorilla/context"
  "stelita-api/wallet"
  "stelita-api/walletCurrency"
)


/**
 *
 */
func GetWalletCurrenciesForWallet(w http.ResponseWriter, r *http.Request) {
  userIdInterface := context.Get(r, "userId")
  var userId int = userIdInterface.(int)

  wallet := wallet.GetUsersWallet(userId)
  walletCurrencies := walletCurrency.GetWalletCurrenciesForWallet(wallet.Id)

  json.NewEncoder(w).Encode(walletCurrencies)
}
