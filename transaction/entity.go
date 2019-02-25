package transaction


type TransactionModel struct {
  Id string
  UserId string
  DepositCurrency string
  DepositAmount float64
  WithdrawalCurrency string
  WithdrawalAmount float64
  ExchangeRate float64
  DateCreated string
}
