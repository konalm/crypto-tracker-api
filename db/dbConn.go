package db

import (
  "os"
  "database/sql"
  "bytes"
  _ "github.com/go-sql-driver/mysql"
)

/**
 * Open connection to database
 */
func Conn() *sql.DB {
  var connBuffer bytes.Buffer
  connBuffer.WriteString(os.Getenv("DB_USER"))
  connBuffer.WriteString(":")
  connBuffer.WriteString(os.Getenv("DB_PASSW"))
  connBuffer.WriteString("@tcp(127.0.0.1)/")
  connBuffer.WriteString(os.Getenv("DB_NAME"))

  // db, err := sql.Open("mysql", "root:$$superstar@tcp(127.0.0.1)/stelita_stag")
  db, err := sql.Open("mysql", connBuffer.String())
  if err != nil {
    panic("ERROR connecting to the database")
    panic(err.Error())
  }

  return db
}
