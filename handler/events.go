package handler

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
  "stelita-api/event"
)


/**
 *
 */
func GetEventReports(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Get Event Items")

  params := mux.Vars(r)
  eventId := params["event_id"]

  fmt.Println("event id >>>")
  fmt.Println(eventId)

  eventReports := event.GetEventReportItems(eventId)

  json.NewEncoder(w).Encode(eventReports)
}
