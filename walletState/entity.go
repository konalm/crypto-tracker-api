package walletState


type WalletState struct {
  Id string
  WalletId string
  WalletStateJson string
  TransactionId string
}

type WalletStateModel struct {
  WalletId string
  UserId int
  WalletDateCreated string
  Id string
  Currencies []WalletStateModelCurrency
  TransactionId string
  DateCreated string
}

type WalletStateModelCurrencies struct {
  Currencies []WalletStateModelCurrency
}

type WalletStateModelCurrency struct {
  Currency string
  Symbol string
  Amount float64
}

type CurrencyInWallet struct {
  Currency string
  Amount float64
}
