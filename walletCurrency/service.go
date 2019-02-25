package walletCurrency

import (
  "fmt"
)


/**
 *
 */
func HandleWalletCurrencyUpdate(walletId string, currency string, amount float64) {
  walletCurrency := GetWalletCurrency(walletId, currency)

  if walletCurrency.Id == "" {
    fmt.Println("wallet currency does NOT exist >> create")

    CreateWalletCurrency(walletId, currency, amount)

  } else {
    fmt.Println("wallet currency DOES exist >> update")

    UpdateWalletCurrencyAmount(walletCurrency.Id, amount)
  }
}
