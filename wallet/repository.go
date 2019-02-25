package wallet

import (
  "fmt"
  "database/sql"
  "os/exec"
  "strings"
  "stelita-api/db"
)


/**
 *
 */
func CreateWallet(userId int) string {
  fmt.Println("create wallet !!")

  uuidOut, err := exec.Command("uuidgen").Output()
  if err != nil {
    panic(err.Error())
  }
  uuid := fmt.Sprintf("%s", uuidOut)
  uuid = strings.TrimSpace(uuid)

  db := db.Conn()
  defer db.Close()

  query := `INSERT INTO wallets (id, user_id) VALUES(?,?)`

  stmt, err := db.Prepare(query)
  if err != nil {
    panic(err.Error())
  }

  _, err = stmt.Exec(uuid, userId)
  if err != nil {
    panic(err.Error())
  }
  defer stmt.Close()


  return uuid
}


/**
 *
 */
func GetUsersWallet(userId int) Wallet {
  db := db.Conn()
  defer db.Close()

  query := "SELECT id, user_id, date_created FROM wallets WHERE user_id = ?"
  stmt := db.QueryRow(query, userId)

  var wallet Wallet
  err := stmt.Scan(&wallet.Id, &wallet.UserId, &wallet.DateCreated)
  if err != nil {
    fmt.Println("ERROR scanning wallet")

    if err == sql.ErrNoRows {
      return Wallet {Id: "", UserId: 0, DateCreated: ""}
    }

    panic(err.Error())
  }

  return wallet
}
