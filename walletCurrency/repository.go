package walletCurrency

import (
  "fmt"
  "os/exec"
  "database/sql"
  "stelita-api/db"
)

/**
 *
 */
func GetWalletCurrency(walletId string, currencyId string) WalletCurrencyModel {
  db := db.Conn()
  defer db.Close()

  query :=
    `SELECT id, wallet_id, currency, amount
    FROM wallet_currencies
    WHERE wallet_id = ?
      AND currency = ?`

  stmt := db.QueryRow(query, walletId, currencyId)

  var walletCurrency WalletCurrencyModel
  err := stmt.Scan(
           &walletCurrency.Id,
           &walletCurrency.WalletId,
           &walletCurrency.Currency,
           &walletCurrency.Amount,
          )
  if err != nil {
    if err == sql.ErrNoRows {
      return WalletCurrencyModel {Id: "", WalletId: "", Currency: "", Amount: 0.00}
    }

    panic(err.Error())
  }

  return walletCurrency
}


/**
 *
 */
func CreateWalletCurrency(walletId string, currencySymbol string, amount float64) {
  uuidOut, err := exec.Command("uuidgen").Output()
  if err != nil {
    panic(err.Error())
  }
  uuid := fmt.Sprintf("%s", uuidOut)

  db := db.Conn()
  defer db.Close()

  query :=
    `INSERT INTO wallet_currencies
    (id, wallet_id, currency, amount)
    VALUES (?,?,?,?)`

  stmt, err := db.Prepare(query)
  if err != nil {
    panic(err.Error())
  }

  _, err = stmt.Exec(uuid, walletId, currencySymbol, amount)
  if err != nil {
    panic(err.Error())
  }
  defer stmt.Close()
}


/**
 *
 */
func UpdateWalletCurrencyAmount(walletCurrencyId string, amount float64) {
  db := db.Conn()
  defer db.Close()

  query := `UPDATE wallet_currencies SET amount = ? WHERE id = ?`

  stmt, err := db.Prepare(query)
  if err != nil {
    panic(err.Error())
  }

  _, err = stmt.Exec(amount, walletCurrencyId)
  if err != nil {
    panic(err.Error())
  }
  defer stmt.Close()
}


/**
 *
 */
func GetWalletCurrencyFromWalletIdCurrency(walletId string, currency string) WalletCurrencyModel {
  fmt.Println("Get wallet currency amount")

  db := db.Conn()
  defer db.Close()

  query :=
    `SELECT id, wallet_id, currency, amount
    FROM wallet_currencies
    WHERE wallet_id = ?
      AND currency = ?
    LIMIT 1`

  stmt := db.QueryRow(query, walletId, currency)

  var walletCurrency WalletCurrencyModel
  err := stmt.Scan(
           &walletCurrency.Id,
           &walletCurrency.WalletId,
           &walletCurrency.Currency,
           &walletCurrency.Amount,
         )
  if err != nil {
    if err == sql.ErrNoRows {
      return WalletCurrencyModel {Id: "", WalletId: "", Currency: "", Amount: 0.00}
    }

    fmt.Println("Error getting wallet currency from wallet id and currency")
    panic(err.Error())
  }

  return walletCurrency
}


/**
 *
 */
func GetWalletCurrenciesForWallet(walletId string) []WalletCurrencyModel {
  db := db.Conn()
  defer db.Close()

  query :=
    `SELECT wc.id, wc.wallet_id, wc.currency, wc.amount,
      IF (ranked_crypto.name IS NULL, "", ranked_crypto.name) name
    FROM wallet_currencies wc
    LEFT JOIN ranked_crypto_currencies ranked_crypto
      ON ranked_crypto.symbol = wc.currency
    WHERE wc.wallet_id = ?
      AND wc.amount > 0.00`

  rows, err := db.Query(query, walletId)
  if err != nil {
    panic("Getting wallet currencies for wallet " + err.Error())
  }
  defer rows.Close()

  var walletCurrencies []WalletCurrencyModel

  for rows.Next() {
    fmt.Println("Get wallet currencies for wallet >> ROW")

    var walletCurrency WalletCurrencyModel
    var currencyName string

    err := rows.Scan(
             &walletCurrency.Id,
             &walletCurrency.WalletId,
             &walletCurrency.Symbol,
             &walletCurrency.Amount,
             &currencyName,
           )
    if err != nil {
      panic("scanning wallet currency for got get wallet currencies for wallet >> " + err.Error())
    }

    fmt.Println("symbol >>")
    fmt.Println(walletCurrency.Symbol)

    if walletCurrency.Symbol == "USD" {
      walletCurrency.Currency = "Dollar"
    } else {
      walletCurrency.Currency = currencyName
    }

    walletCurrencies = append(walletCurrencies, walletCurrency)
  }

  return walletCurrencies
}
