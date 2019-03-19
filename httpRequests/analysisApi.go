package httpRequests

import (
  "net/http"
  "fmt"
  "io/ioutil"
)


/**
 *
 */
func UpdateAnalysisReports() {
  fmt.Println("http requests >> update analysis reports")

  response, err := http.Get("http://138.68.167.173:3001/update-analysis-report")
  if err != nil {
    fmt.Println(err)
  }

  defer response.Body.Close()
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(body)
}
