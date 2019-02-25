package handler

import (
  "fmt"
  "net/http"
  "github.com/gorilla/context"
  "encoding/json"
  "stelita-api/wallet"
  "stelita-api/deposit"
  "stelita-api/transaction"
  "stelita-api/walletState"
  "stelita-api/walletCurrency"
)


/**
 *
 */
func SetupWallet (w http.ResponseWriter, r *http.Request) {
  fmt.Println("Setup Wallet")

  userIdInterface := context.Get(r, "userId")
  var userId int = userIdInterface.(int)

  usersWallet := wallet.GetUsersWallet(userId)

  /* check if user already has a wallet */
  if usersWallet.Id != "" {
    w.WriteHeader(403)
    w.Write([]byte("user already has wallet setup"))
    return
  }

  deposit := deposit.CreateDeposit("USD", 10000.00)
  transaction := transaction.CreateTransactionModel(userId, "", deposit, 1.00)
  wallet := wallet.CreateWallet(userId)

  walletCurrency.HandleWalletCurrencyUpdate(
    wallet,
    "USD",
    10000.00,
  )

  walletState.CreateWalletState(wallet, transaction)

  w.Write([]byte("setup wallet endpoint"))
}


/**
 *
 */
func GetCurrenciesInWallet (w http.ResponseWriter, r *http.Request) {
  fmt.Println("Get currencies in wallet")

  userIdInterface := context.Get(r, "userId")
  var userId int = userIdInterface.(int)

  usersWallet := wallet.GetUsersWallet(userId)
  latestWalletState := walletState.GetWalletsLatestState(usersWallet.Id)

  currencies := walletState.GetCurrenciesFromWalletState(latestWalletState.Currencies)

  json.NewEncoder(w).Encode(currencies)
}
