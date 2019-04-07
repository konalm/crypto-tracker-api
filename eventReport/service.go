package eventReport

import (
  "fmt"
  "stelita-api/errorReporter"
)


/**
 *
 */
func ReportEvent(event string, message string, error bool) {
  fmt.Println("Service >> report event")

  if error {
    fmt.Println("is error >> report to error reporter")
    errorReporter.ReportError(message)
  }

  InsertEventReport(event, message, error)
}
