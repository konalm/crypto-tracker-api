package httpRequests

import (
  "net/http"
  "fmt"
  "strconv"
  "stelita-api/config"
  "stelita-api/eventReport"
  "stelita-api/logger"
)


/**
 *

 */
func UpdateAnalysisReports() {
  logger.WriteLog("update analysis reports HTTP REQUEST", "")

  response, err := http.Get(config.ANALYSIS_API_URL + "/update-analysis-report")
  if err != nil {
    eventReport.ReportEvent("21e90f38-3a6c-4599-9a22-2fb5de69c05b",
      "update analysis reports http request, could not connect to Analysis API",
      true,
    )
    logger.WriteLog("update analysis reports http request, could not connect to Analysis API", "")
    return
  }

  if response.StatusCode != 200 {
    statusCodeString := strconv.Itoa(response.StatusCode)
    eventReport.ReportEvent("21e90f38-3a6c-4599-9a22-2fb5de69c05b",
      "Http request to analysis reports resulted in status code " + statusCodeString,
      true,
    )
    logger.WriteLog("Http request to analysis reports resulted in status code " + statusCodeString, "")
    return
  }
  defer response.Body.Close()

  eventReport.ReportEvent("21e90f38-3a6c-4599-9a22-2fb5de69c05b", "", false)
}


/**
 *
 */
func StartAnalysisReports() {
  logger.WriteLog("start analysis reports HTTP REQUEST", "")

  response, err := http.Get(config.ANALYSIS_API_URL + "/start-analysis-cryptos-to-analyse")
  if err != nil {
    fmt.Println("start analysis reports htttp request, could not connect to Analysis API")
    eventReport.ReportEvent("2da0da5a-2872-4528-bd8f-e29eced4df22",
      "start analysis reports htttp request, could not connect to Analysis API",
      true,
    )
    return
  }

  if response.StatusCode != 200 {
    eventReport.ReportEvent("2da0da5a-2872-4528-bd8f-e29eced4df22",
      "Http request to start analysis reports resulted in status code " + strconv.Itoa(response.StatusCode),
      true,
    )
  } else {
    eventReport.ReportEvent("2da0da5a-2872-4528-bd8f-e29eced4df22", "", false)
  }

  defer response.Body.Close()
}
