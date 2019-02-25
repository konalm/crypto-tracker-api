package handler

import (
  "fmt"
  "net/http"
  "github.com/gorilla/context"
  "encoding/json"
  "stelita-api/walletState"
  "stelita-api/wallet"
)


/**
 *
 */
func GetUserWalletStates(w http.ResponseWriter, r *http.Request) {
  fmt.Println("get user wallet states !!")

  userIdInterface := context.Get(r, "userId")
  var userId int = userIdInterface.(int)

  walletStates := walletState.GetUserWalletStates(userId)

  json.NewEncoder(w).Encode(walletStates)
}


/**
 *
 */
func GetUserLatestWalletState(w http.ResponseWriter, r *http.Request) {
  fmt.Println("get user latest wallet state")

  userIdInterface := context.Get(r, "userId")
  var userId int = userIdInterface.(int)

  usersWallet := wallet.GetUsersWallet(userId)
  latestWalletState := walletState.GetWalletsLatestState(usersWallet.Id)

  json.NewEncoder(w).Encode(latestWalletState)
}
