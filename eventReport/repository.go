package eventReport

import (
  "fmt"
  "stelita-api/db"
  "stelita-api/errorReporter"
)

/**
 *
 */
func InsertEventReport(event string, description string, success bool) {
  fmt.Println("insert event report")
  fmt.Println(event)
  fmt.Println(description)
  fmt.Println(success)
  fmt.Println("<<<<<<<<<<<<<<<<")

  conn := db.Conn()
  defer conn.Close()

  query :=
    `INSERT INTO event_reports
    (event, description, success)
    VALUES (?,?,?)`

  stmt, err := conn.Prepare(query)
  if err != nil {
    errorReporter.ReportError("failed to prepare insert statement for event report")
    panic(err.Error())
  }
  defer stmt.Close()

  _, err = stmt.Exec(event, description, success)
  if err != nil {
    errorReporter.ReportError("failed to execute insert for event report")
    panic(err.Error())
  }
}
