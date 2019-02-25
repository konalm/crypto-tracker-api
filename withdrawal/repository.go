package withdrawal

import (
  "fmt"
  "os/exec"
  "stelita-api/db"
)

/**
 *
 */
func CreateWithdrawal(cryptoCurrency string, amount float64) string {
  fmt.Println("create withdrawal")

  uuidOut, err := exec.Command("uuidgen").Output()
  if err != nil {
    panic(err.Error())
  }
  uuid := fmt.Sprintf("%s", uuidOut)

  db := db.Conn()
  defer db.Close()

  query := `INSERT INTO withdrawals (id, crypto_currency, amount) VALUES(?,?,?)`

  stmt, err := db.Prepare(query)
  if err != nil {
    panic(err.Error())
  }

  _, err = stmt.Exec(uuid, cryptoCurrency, amount)
  if err != nil {
    panic(err.Error())
  }

  defer stmt.Close()

  return uuid
}
