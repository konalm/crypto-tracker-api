package event

import (
  "fmt"
  "stelita-api/db"
)

/**
 *
 */
func GetEventReportItems(eventId string) []EventReport {
  fmt.Println("Report >> get event report items")

  conn := db.Conn()
  defer conn.Close()

  query :=
    `SELECT e.event AS event_name, er.description, er.success, er.date_created
    FROM events e
    INNER JOIN event_reports er
      ON er.event = e.id`
  queryValues := []interface{}{}

  if eventId != "*" {
    query = query + " WHERE e.id = ?"
    queryValues = append(queryValues, eventId)
  }

  query = query + " ORDER BY er.date_created DESC LIMIT 250"

  rows, err := conn.Query(query, queryValues...)
  if err != nil {
    fmt.Println("ERROR executing query")
    panic(err.Error())
  }
  defer rows.Close()

  var eventReports []EventReport

  for rows.Next() {
    var eventReport EventReport

    err := rows.Scan(
      &eventReport.EventName,
      &eventReport.Description,
      &eventReport.Success,
      &eventReport.DateCreated,
    )
    if err != nil {
      panic(err.Error())
    }

    eventReports = append(eventReports, eventReport)
  }

  return eventReports
}
