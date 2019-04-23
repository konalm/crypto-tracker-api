package cryptoRatesController

import (
	// "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
	"stelita-api/structs"
	"stelita-api/db"
	"stelita-api/reports"
	"stelita-api/errorReporter"
	"stelita-api/fetchCryptoRates"
)


/**
 *
 */
func InsertCryptoRates(rates []structs.BitcoinRate) {
	fmt.Println("Insert Crypto Rates")

  dbConn := db.Conn()
	defer dbConn.Close()

  query := `INSERT INTO crypto_rates (currency, date, closing_price, min)`
  var queryValues = []interface{}{}

  timeOfCronJob := time.Now()
  minOfCronJob := timeOfCronJob.Minute()

  for i, rate := range rates {
    query += buildRateModalQuery(i)
    formattedDate := formatDateForMysql(rate.Time)

    if rate.Asset_id_quote == "USD" { rate.Asset_id_quote = "BTC" }

    queryValues = append(queryValues, rate.Asset_id_quote, formattedDate, rate.Rate, minOfCronJob)
  }

  query = removeLastComma(query)
  stmt, _ := dbConn.Prepare(query)
	defer stmt.Close()

	insertSuccess := true
  _, err := stmt.Exec(queryValues...)
  if err != nil {
		errorReporter.ReportError("Inserting crypo rates")
    insertSuccess = false
  }

	processListQuery := "SHOW PROCESSLIST";
	stmtProcessList, err := dbConn.Query(processListQuery)
	if err != nil {
		errorReporter.ReportError("Executing process list on insert crypto rates")
		fmt.Println("stmt process list")
	}
	defer stmtProcessList.Close()

	i := 0
	for stmtProcessList.Next() {
		i ++
	}

  reports.InsertCryptoReport("cryptoRates", insertSuccess, i)
}

/**
 *
 */
func InsertCryptoRatesFromCoinMarketCapApi(rates []fetchCryptoRates.CoinMarketCapCryptoCurrency) {
	fmt.Println("Insert crypto rates from coin market cap api")

  dbConn := db.Conn()
	defer dbConn.Close()

  query := `INSERT INTO crypto_rates (currency, date, closing_price, min)`
  var queryValues = []interface{}{}

  timeOfCronJob := time.Now()
  minOfCronJob := timeOfCronJob.Minute()

  for i, rate := range rates {
    query += buildRateModalQuery(i)
    formattedDate := formatDateForMysql(rate.Last_updated)

    queryValues = append(queryValues,
			rate.Symbol,
			formattedDate,
			rate.Quote.BTC.Price,
			minOfCronJob,
		)
  }

  query = removeLastComma(query)
  stmt, _ := dbConn.Prepare(query)
	defer stmt.Close()

	insertSuccess := true
  _, err := stmt.Exec(queryValues...)
  if err != nil {
		errorReporter.ReportError("Inserting crypo rates")
    insertSuccess = false
  }

	processListQuery := "SHOW PROCESSLIST";
	stmtProcessList, err := dbConn.Query(processListQuery)
	if err != nil {
		errorReporter.ReportError("Executing process list on insert crypto rates")
		fmt.Println("stmt process list")
	}
	defer stmtProcessList.Close()

	i := 0
	for stmtProcessList.Next() {
		i ++
	}

  reports.InsertCryptoReport("cryptoRates", insertSuccess, i)
}

/**
 *
 */
func InsertUSDCryptoRates(rates []structs.USDRate) {
	fmt.Println("INSERT USD CRYPTO RATES")

	dbConn := db.Conn()
	defer dbConn.Close()

	query := `INSERT crypto_usd_rates (current, date, closing_price, min)`
	var queryValues = []interface{}{}

	timeOfCronJob := time.Now()
	minOfCronJob := timeOfCronJob.Minute()

	for i, rate := range rates {
		query += buildRateModalQuery(i)
		formattedDate := formatDateForMysql(rate.Date_added)

		queryValues =
			append(queryValues, rate.Symbol, formattedDate, rate.Quotes[0].Price, minOfCronJob)

		query = removeLastComma(query)
		stmt, _ := dbConn.Prepare(query)
		defer stmt.Close()

		_, err := stmt.Exec(queryValues...)
		if err != nil {
			errorReporter.ReportError("Inserting USD crypto rates")
			panic(err.Error())
		}

		processListQuery := "SHOW PROCESSLIST";
		stmtProcessList, err := dbConn.Query(processListQuery)
		if err != nil {
			fmt.Println("stmt process list")
		}
		defer stmtProcessList.Close()

		i := 0
		for stmtProcessList.Next() {
			i ++
		}

		fmt.Println("Process List -- on insert -- >>>>>>>")
		fmt.Println(i)
	}
}


/**
 * build query to insert single rate modal
 */
func buildRateModalQuery(i int) string {
  var query string

  if i == 0 {
    query += " VALUES"
  }

  query += " (?,?,?,?),"

  return query
}


/**
 *
 */
func formatDateForMysql(date string) time.Time {
  var replacer = strings.NewReplacer("T", " ", "Z", "")
  preparedDate := replacer.Replace(date)

  formattedDate, err := time.Parse("2006-01-02 15:04:05", preparedDate)
  if err != nil {
    fmt.Println("Error: could not format date")
  }

  return formattedDate
}


/**
 * Remove the last comma from the query string
 */
func removeLastComma(query string) string  {
  queryx := len(query)

  if queryx > 0 && query[queryx-1] == ',' {
    query = query[:queryx-1]
  }

  return query
}
