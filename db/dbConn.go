package db

import (
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

/**
 * Open connection to database
 */
func Conn() *sql.DB {
  fmt.Println("open connection to database")

  // db, err := sql.Open("mysql", "root:$$superstar@tcp(127.0.0.1:5432)/stelita_stag")
  db, err := sql.Open("mysql", "root:$$superstar@tcp(127.0.0.1)/stelita_stag")
  if err != nil {
    panic("ERROR connecting to the database")
    panic(err.Error())
  }

  return db
}
