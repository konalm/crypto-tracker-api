package reports

import (
  "fmt"
  "encoding/json"
  "net/http"
  "stelita-api/db"
)

/**
 * Get reports on Crypto Data
 */
func GetCryptoDataReports(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Get crypto data reports !!")

  db := db.Conn()
  defer db.Close()

  type CryptoDataReport struct {
    RequestStatus int
    DateCreated string
  }

  query :=
    `SELECT request_status, date_created
    FROM fetch_crypto_data_reports
    ORDER BY date_created DESC`

  rows, err := db.Query(query)
  if err != nil {
    panic(err.Error())
  }
  defer rows.Close()

  var cryptoDataReports []CryptoDataReport

  for rows.Next() {
    var cryptoDataReport CryptoDataReport

    err := rows.Scan(
             &cryptoDataReport.RequestStatus,
             &cryptoDataReport.DateCreated,
           )
    if err != nil {
      panic(err.Error())
    }

    cryptoDataReports = append(cryptoDataReports, cryptoDataReport)
  }

  json.NewEncoder(w).Encode(cryptoDataReports)
}


/**
 * Get Reports on Ranked Crypto Currencies
 */
func GetRankedCryptoCurrenciesReports(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Get ranked crypto currencies reports !!")

  db := db.Conn()
  defer db.Close()

  type RankedCryptoCurrenciesReport struct {
    RequestStatus int
    DateCreated string
  }

  query :=
    `SELECT request_status, date_created
    FROM fetch_ranked_crypto_currencies_reports
    ORDER BY date_created DESC`

  rows, err := db.Query(query)
  if err != nil {
    panic(err.Error())
  }
  defer rows.Close()

  var rankedCryptoCurrenciesReports []RankedCryptoCurrenciesReport

  for rows.Next() {
    var rankedCryptoCurrenciesReport RankedCryptoCurrenciesReport

    err := rows.Scan(
             &rankedCryptoCurrenciesReport.RequestStatus,
             &rankedCryptoCurrenciesReport.DateCreated,
           )
    if err != nil {
      panic(err.Error())
    }

    rankedCryptoCurrenciesReports =
      append(rankedCryptoCurrenciesReports, rankedCryptoCurrenciesReport)
  }

  json.NewEncoder(w).Encode(rankedCryptoCurrenciesReports)
}


/**
 * Get Reports on Inserted Crypto
 */
func GetInsertedCryptoReports(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Get insert crypto reports")

  db := db.Conn()
  defer db.Close()

  type InsertCryptoReport struct {
    DateCreated string
    Success int
    DBProcessList int
  }

  query :=
    `SELECT date_created, success, db_process_list
    FROM insert_crypto_reports
    ORDER BY date_created DESC`

  rows, err := db.Query(query)
  if err != nil {
    panic(err.Error())
  }
  defer rows.Close()

  var insertCryptoReports []InsertCryptoReport

  for rows.Next() {
    var insertCryptoReport InsertCryptoReport

    err := rows.Scan(
             &insertCryptoReport.DateCreated,
             &insertCryptoReport.Success,
             &insertCryptoReport.DBProcessList,
           )
    if err != nil {
      panic(err.Error())
    }

    insertCryptoReports = append(insertCryptoReports, insertCryptoReport)
  }

  json.NewEncoder(w).Encode(insertCryptoReports)
}


/**
 *
 */
func GetUpdateCryptoTrendStatReports(w http.ResponseWriter, r *http.Request) {
  fmt.Println("get updated crypto trend stat reports")

  db := db.Conn()
  defer db.Close()

  type Report struct {
    CryptoCurrency string
    DateCreated string
    Success int
    DBProcessList int
  }

  query :=
    `SELECT crypto_currency, date_created, success, db_process_list
    FROM update_crypto_trend_stat_reports
    ORDER BY date_created DESC`

  rows, err := db.Query(query)
  if err != nil {
    panic(err.Error())
  }
  defer rows.Close()

  var reports []Report

  for rows.Next() {
    var report Report

    err := rows.Scan(
             &report.CryptoCurrency,
             &report.DateCreated,
             &report.Success,
             &report.DBProcessList,
           )
    if err != nil {
      panic(err.Error())
    }

    reports = append(reports, report)
  }

  json.NewEncoder(w).Encode(reports)
}
