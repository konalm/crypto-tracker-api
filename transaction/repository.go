package transaction


import (
  "fmt"
  "strings"
  "database/sql"
  "os/exec"
  "stelita-api/db"
)


/**
 *
 */
func GetTransactionModel(_transactionId string) TransactionModel {
  transactionId := strings.TrimSpace(_transactionId)

  db := db.Conn()
  defer db.Close()

  query :=
    `SELECT t.id, t.user_id, t.exchange_rate, t.date_created,
      d.crypto_currency AS deposit_currency, d.amount deposit_amount,
      IF (w.crypto_currency IS NULL, "", w.crypto_currency) withdrawal_currency,
      IF (w.amount IS NULL, 0.00, w.amount) withdrawal_amount
    FROM transactions t
    LEFT JOIN deposits d
      ON d.id = t.deposit_id
    LEFT JOIN withdrawals w
      ON w.id = t.withdrawal_id
    WHERE t.id = ?`

  stmt := db.QueryRow(query, transactionId)

  var transactionModel TransactionModel
  err := stmt.Scan(
           &transactionModel.Id,
           &transactionModel.UserId,
           &transactionModel.ExchangeRate,
           &transactionModel.DateCreated,
           &transactionModel.DepositCurrency,
           &transactionModel.DepositAmount,
           &transactionModel.WithdrawalCurrency,
           &transactionModel.WithdrawalAmount,
         )
  if err != nil {
    if err == sql.ErrNoRows {
      return TransactionModel{
               Id: "",
               UserId: "",
               DepositCurrency: "",
               DepositAmount: 0.00,
               WithdrawalCurrency: "",
               WithdrawalAmount: 0.00,
               ExchangeRate: 0.00,
               DateCreated: "",
             }
    }

    panic(err.Error())
  }

  return transactionModel
}


/**
 *
 */
func CreateTransactionModel(
  userId int, withdrawalId string, depositId string, exchangeRate float64,
) string {
  uuidOut, err := exec.Command("uuidgen").Output()
  if err != nil {
    panic(err.Error())
  }
  uuid := fmt.Sprintf("%s", uuidOut)

  db := db.Conn()
  defer db.Close()

  query :=
    `INSERT INTO transactions
    (id, user_id, withdrawal_id, deposit_id, exchange_rate)
    VALUES (?,?,?,?,?)`

  stmt, err := db.Prepare(query)
  if err != nil {
    panic(err.Error())
  }

  _, err = stmt.Exec(uuid, userId, withdrawalId, depositId, exchangeRate)
  if err != nil {
    panic(err.Error())
  }
  defer stmt.Close()

  return uuid
}


/**
 *
 */
func GetUserTransactions(userId int) []TransactionModel {
  db := db.Conn()
  defer db.Close()

  query :=
    `SELECT t.id, t.user_id, t.exchange_rate, t.date_created,
      IF (w.crypto_currency IS NULL, "", w.crypto_currency) withdrawal_currency,
      IF (w.amount IS NULL, 0.00, w.amount) withdrawal_amount,
      d.crypto_currency deposit_currency, d.amount deposit_amount
    FROM transactions t
    LEFT JOIN deposits d
      ON d.id = t.deposit_id
    LEFT JOIN withdrawals w
      ON w.id = t.withdrawal_id
    WHERE t.user_id = ?
    ORDER BY t.date_created DESC`

  rows, err := db.Query(query, userId)
  if err != nil {
    panic(err.Error())
  }
  defer rows.Close()

  var transactionModels []TransactionModel

  for rows.Next() {
    var transactionModel TransactionModel

    err := rows.Scan(
             &transactionModel.Id,
             &transactionModel.UserId,
             &transactionModel.ExchangeRate,
             &transactionModel.DateCreated,
             &transactionModel.WithdrawalCurrency,
             &transactionModel.WithdrawalAmount,
             &transactionModel.DepositCurrency,
             &transactionModel.DepositAmount,
           )
    if err != nil {
      panic(err.Error())
    }

    transactionModels = append(transactionModels, transactionModel)
  }

  return transactionModels
}
