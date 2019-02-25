package walletState

import (
  "fmt"
  "database/sql"
  "os/exec"
  "encoding/json"
  "stelita-api/db"
)

/**
 *
 */
func GetUserWalletStates(userId int) []WalletStateModel {
  db := db.Conn()
  defer db.Close()

  query :=
    `SELECT w.id wallet_id, w.user_id, w.date_created wallet_date_created,
      ws.id, ws.wallet_state_json, ws.transaction_id, ws.date_created
    FROM wallets w
    LEFT JOIN wallet_states ws
      ON ws.wallet_id = w.id
    WHERE w.user_id = ?
    ORDER BY ws.date_created DESC`

  rows, err := db.Query(query, userId)
  if err != nil {
    panic(err.Error())
  }
  defer rows.Close()

  var walletStates []WalletStateModel

  for rows.Next() {
    var walletState WalletStateModel
    var walletStateCurrenciesJson string

    err := rows.Scan(
             &walletState.WalletId,
             &walletState.UserId,
             &walletState.DateCreated,
             &walletState.Id,
             &walletStateCurrenciesJson,
             &walletState.TransactionId,
             &walletState.DateCreated,
           )
    if err != nil {
      panic(err.Error())
    }

    var walletStateModelCurrencies WalletStateModelCurrencies
    json.Unmarshal([]byte(walletStateCurrenciesJson), &walletStateModelCurrencies)
    walletState.Currencies = walletStateModelCurrencies.Currencies

    walletStates = append(walletStates, walletState)
  }

  return walletStates
}


/**
 * Create new wallet state model
 */
func CreateWalletState (walletId string, transactionId string) {
  uuidOut, err := exec.Command("uuidgen").Output()
  if err != nil {
    panic("getting user id >> " + err.Error())
  }
  uuid := fmt.Sprintf("%s", uuidOut)

  db := db.Conn()
  defer db.Close()

  // walletStateJson := CalcWalletState(walletId, transactionId)
  walletStateJson := BuildWalletStateJson(walletId)

  query :=
    `INSERT wallet_states
    (id, wallet_id, wallet_state_json, transaction_id)
    VALUES(?,?,?,?)`

  stmt, err := db.Prepare(query)
  if err != nil {
    panic("prepare wallet status >> " + err.Error())
  }

  _, err = stmt.Exec(uuid, walletId, walletStateJson, transactionId)
  if err != nil {
    panic("insert wallet status >> " + err.Error())
  }
}


/**
 * Get last created wallet state model related to wallet
 */
func GetWalletsLatestState(walletId string) WalletStateModel {
  db := db.Conn()
  defer db.Close()

  query :=
    `SELECT wallet_states.id, wallet_state_json, wallet_states.transaction_id,
      wallet_states.date_created,
      wallets.id wallet_id, wallets.user_id, wallets.date_created wallet_date_created
    FROM wallet_states
    INNER JOIN wallets
      ON wallets.id = wallet_states.wallet_id
    WHERE wallet_id = ?
    ORDER BY date_created DESC
    LIMIT 1`
  stmt := db.QueryRow(query, walletId)

  var walletState WalletStateModel
  var walletStateCurrenciesJson string

  err := stmt.Scan(
           &walletState.Id,
           &walletStateCurrenciesJson,
           &walletState.TransactionId,
           &walletState.DateCreated,
           &walletState.WalletId,
           &walletState.UserId,
           &walletState.WalletDateCreated,
         )
  if err != nil {
    if err == sql.ErrNoRows {
      return WalletStateModel{Id: "", WalletId: "", TransactionId: ""}
    }

    fmt.Println("Error getting wallets latest state")
    panic(err.Error())
  }

  var walletStateModelCurrencies WalletStateModelCurrencies
  json.Unmarshal([]byte(walletStateCurrenciesJson), &walletStateModelCurrencies)
  walletState.Currencies = walletStateModelCurrencies.Currencies

  return walletState
}
