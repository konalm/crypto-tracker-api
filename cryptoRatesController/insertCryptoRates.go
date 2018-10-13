package cryptoRatesController

import (
	"crypto-tracker-api/structs"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

type CryptoRate struct {
	currency     string
	date         time.Time
	closingPrice float64
	min          int
}

func InsertCryptoRates(rates []structs.BitcoinRate) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/stelita_dev")
	if err != nil {
		panic(err.Error())
	}

	var queryValues []CryptoRate

	query := "INSERT INTO crypto_rates (currency, date, closing_price, min)"
	for _, rate := range rates {
		var replacer = strings.NewReplacer("T", " ", "Z", "")
		preparedDate := replacer.Replace(rate.Time)
		formattedDate, err := time.Parse("2006-01-02 15:04:05", preparedDate)
		if err != nil {
			fmt.Println(err.Error())
		}
		minOfCronJob := time.Now().Minute()
		query += " VALUES(?,?,?,?),"
		queryValues = append(queryValues, CryptoRate{
			rate.Asset_id_quote,
			formattedDate,
			rate.Rate,
			minOfCronJob,
		})
	}

	lq := len(query)
	if lq > 0 && query[lq-1] == ',' {
		query = query[:lq-1]
	}

	fmt.Println("Query >>>>")
	fmt.Println(query)

	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(queryValues)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()
	defer db.Close()
}
