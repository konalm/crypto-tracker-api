package indicatorReporter

import (
  "fmt"
  "stelita-api/wallet"
  "stelita-api/walletCurrency"
  "stelita-api/rankedCryptoCurrency"
  "stelita-api/email"
)


func ReportIndicatorsViaEmail(userId int) {
  fmt.Println("report indicators via email")

  wallet :=  wallet.GetUsersWallet(userId)
  currenciesInWallet := walletCurrency.GetWalletCurrenciesForWallet(wallet.Id)
  cryptoCurrencyData := rankedCryptoCurrency.GetCryptoCurrencyData(currenciesInWallet)

  for _, crypto := range cryptoCurrencyData {
    if crypto.TrendStats == nil {
      continue
    }

    rsi := crypto.TrendStats[0].Rsi

    if !crypto.InWallet && crypto.BuyIndicator {
      fmt.Println("crypto not in wallet and ready to buy >>> TRIGGER INDICATOR")
      email.TradeCurrencyNotification(crypto.Name, rsi, "buying")
    }

    if crypto.InWallet && crypto.SellIndicator {
      fmt.Println("crypto in wallet and ready for sale >>> TRIGGER INDICATOR")
      email.TradeCurrencyNotification(crypto.Name, rsi, "selling")
    }
  }
}
