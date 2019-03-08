package walletState

import (
  "fmt"
  "encoding/json"
  "stelita-api/transaction"
  "stelita-api/walletCurrency"
)

/**
 *
 */
func BuildWalletStateJson(walletId string) string {
  var walletState WalletStateModelCurrencies

  walletCurrencies := walletCurrency.GetWalletCurrenciesForWallet(walletId)

  for _, currency := range walletCurrencies {
    fmt.Println("in currency loop")

    var walletStateModelCurrency = walletCurrency.WalletCurrencyModel {
      Currency: currency.Currency,
      Symbol: currency.Symbol,
      Amount: currency.Amount,
    }

    walletState.Currencies = append(walletState.Currencies, walletStateModelCurrency)
  }

  walletStateJson, err := json.Marshal(walletState)
  if err != nil {
    panic(err.Error())
  }

  return string(walletStateJson)
}


/**
 *
 */
func CalcWalletState(walletId string, transactionId string) string {
  type WalletState struct {
    Currencies []CurrencyInWallet
  }

  lastWalletState := GetWalletsLatestState(walletId)

  var newWalletState WalletStateModelCurrencies

  if lastWalletState.Id == "" {
    var currencyInWallet = walletCurrency.WalletCurrencyModel {
                             Currency: "Dollar",
                             Symbol: "USD",
                             Amount: 10000.00,
                           }
    newWalletState.Currencies = append(newWalletState.Currencies, currencyInWallet)
  } else {
    /* calc based on previous wallet and transaction */
    transaction := transaction.GetTransactionModel(transactionId)

    walletStateHasCurrency := CheckWalletStateContainsCurrency(
                                lastWalletState.Currencies,
                                transaction.DepositCurrency,
                              )
    newWalletState.Currencies = lastWalletState.Currencies

    if !walletStateHasCurrency {
      var depositCurrency = walletCurrency.WalletCurrencyModel {
                              Currency: transaction.DepositCurrency,
                              Amount: transaction.DepositAmount,
                            }
      newWalletState.Currencies = append(newWalletState.Currencies, depositCurrency)
    } else {
      depositIndex := getIndexOfCurrencyInWallet(
                        lastWalletState.Currencies,
                        transaction.DepositCurrency,
                      )

      newWalletState.Currencies[depositIndex].Amount += transaction.DepositAmount
    }

    withdrawIndex := getIndexOfCurrencyInWallet(
                       lastWalletState.Currencies,
                       transaction.WithdrawalCurrency,
                     )
    newWalletState.Currencies[withdrawIndex].Amount -= transaction.WithdrawalAmount
  }

  walletStateJson, err := json.Marshal(newWalletState)
  if err != nil {
    panic(err.Error())
  }

  return string(walletStateJson)
}


/**
 *
 */
func CheckWalletStateContainsCurrency(
  walletStateCurrencies []walletCurrency.WalletCurrencyModel,
  depositCurrency string,
) bool {
  for _, currency := range walletStateCurrencies {
    fmt.Println("range >>")
    fmt.Println("symbol >>> " + currency.Symbol)
    fmt.Println("currency >>> " + currency.Currency + "!!!")
    fmt.Println("deposit currency >>> " + depositCurrency + "!!!")
    fmt.Println(" <-- -- ")

    if currency.Currency == depositCurrency {
      return true
    }
  }

  return false
}


/**
 *
 */
func getIndexOfCurrencyInWallet(
  walletStateCurrencies []walletCurrency.WalletCurrencyModel,
  depositCurrency string,
) int {
  for i, currency := range walletStateCurrencies {
    if currency.Currency == depositCurrency {
      return i
    }
  }

  return -1
}

/**
 *
 */
func GetCurrenciesFromWalletState(
  walletCurrencies []walletCurrency.WalletCurrencyModel,
) []string {
  var currencies []string

  for _, currency := range walletCurrencies {
    currencies = append(currencies, currency.Currency)
  }

  return currencies
}
