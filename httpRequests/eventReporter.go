package httpRequests

import (
  "net/http"
  "fmt"
  "encoding/json"
  "bytes"
  "stelita-api/config"
)

/**
 *
 */
func ReportError(eventId string, description string, success bool) {
  fmt.Println("REPORT ERROR !!")

  type Event struct {
    Event string
    Description string
    Success bool
  }

  var event = Event{
    Event: eventId,
    Description: description,
    Success: success,
  }
  eventJson, _ := json.Marshal(event)

  fmt.Println("report error")
  fmt.Println(event)
  fmt.Println(description)
  fmt.Println(success)
  fmt.Println("<<<<<<<<<<<<<<")

  request, err := http.NewRequest("POST", config.EVENT_REPORTER_API_URL, bytes.NewBuffer(eventJson))
  request.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  response, err := client.Do(request)
  if err != nil {
    fmt.Println("Error: sending http post request to Event Reporter API")
  }
  defer response.Body.Close()

  if response.StatusCode != 200 {
    fmt.Println("Error: post request to Event Reporter API was not successful")
  }
}
