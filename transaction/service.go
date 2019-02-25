package transaction

import (
  "fmt"
  // "stelita-api/db"
)

/**
 *
 */
func TransactionModelValidation(
  withdrawCurrency string, depositCurrency string, amount float64,
) string {
  fmt.Println("transaction model validation")

  if withdrawCurrency == "" {
    return "withdraw currency is required"
  }

  if depositCurrency == "" {
    return "deposit currency is required"
  }

  if amount == 0 {
    return "amount is required"
  }

  return ""
}
