package handler

import (
  "fmt"
  "net/http"
  "github.com/gorilla/context"
  "strconv"
  "encoding/json"
  "stelita-api/transaction"
  "stelita-api/rankedCryptoCurrency"
  "stelita-api/deposit"
  "stelita-api/withdrawal"
  "stelita-api/wallet"
  "stelita-api/walletState"
  "stelita-api/walletCurrency"
)


/**
 *
 */
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Create transaction !!")

  userIdInterface := context.Get(r, "userId")
  var userId int = userIdInterface.(int)

  withdrawCurrency := r.FormValue("withdrawCurrency")
  depositCurrency := r.FormValue("depositCurrency")

  amountString := r.FormValue("amount")
  amount, err := strconv.ParseFloat(amountString, 64)
  if err != nil {
    panic(err.Error())
  }

  validation := transaction.TransactionModelValidation(withdrawCurrency, depositCurrency, amount)
  if validation != "" {
    w.WriteHeader(406)
    w.Write([]byte(validation))
    return
  }

  exchangeRate, err := rankedCryptoCurrency.FetchCryptoCurrencyExchangeRate(withdrawCurrency, depositCurrency)
  if err != nil {
    w.WriteHeader(550)
    w.Write([]byte(err.Error()))
    return
  }

  withdrawal := withdrawal.CreateWithdrawal(withdrawCurrency, amount)

  depositAmount := amount * exchangeRate
  deposit := deposit.CreateDeposit(depositCurrency, depositAmount)

  transactionId := transaction.CreateTransactionModel(userId, withdrawal, deposit, exchangeRate)

  wallet := wallet.GetUsersWallet(userId)

  walletCurrency.HandleWalletCurrencyUpdate(wallet.Id, depositCurrency, depositAmount)

  withdrawWalletCurrency :=
    walletCurrency.GetWalletCurrencyFromWalletIdCurrency(wallet.Id, withdrawCurrency)
  withdrawWalletCurrencyAmount := withdrawWalletCurrency.Amount - amount

  walletCurrency.HandleWalletCurrencyUpdate(
    wallet.Id,
    withdrawCurrency,
    withdrawWalletCurrencyAmount,
  )

  walletState.CreateWalletState(wallet.Id, transactionId)

  w.Write([]byte("Create transaction endpoint"))
}


/**
 *
 */
func GetUserTransactions(w http.ResponseWriter, r *http.Request) {
  userIdInterface := context.Get(r, "userId")
  var userId int = userIdInterface.(int)

  transactions := transaction.GetUserTransactions(userId)

  json.NewEncoder(w).Encode(transactions)
}
