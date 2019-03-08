package indicatorReporter

import (
  "stelita-api/wallet"
  "stelita-api/walletCurrency"
  "stelita-api/rankedCryptoCurrency"
  "stelita-api/email"
)


func ReportIndicatorsViaEmail(userId int) {
  wallet :=  wallet.GetUsersWallet(userId)
  currenciesInWallet := walletCurrency.GetWalletCurrenciesForWallet(wallet.Id)
  cryptoCurrencyData := rankedCryptoCurrency.GetCryptoCurrencyData(currenciesInWallet)

  for _, crypto := range cryptoCurrencyData {
    if crypto.TrendStats == nil {
      continue
    }

    rsi := crypto.TrendStats[0].RsiStats.Rsi

    if !crypto.InWallet && crypto.BuyIndicator {
      email.TradeCurrencyNotification(crypto.Name, rsi, "buying")
    }

    if crypto.InWallet && crypto.SellIndicator {
      email.TradeCurrencyNotification(crypto.Name, rsi, "selling")
    }
  }
}
