package db

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "stelita-api/structs"
)

func GetProcessList() []structs.ProcessList {
    /* open database connection */
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/stelita_dev")
    if err != nil {
      panic(err.Error())
    }
    defer db.Close()

    query := `SHOW PROCESSLIST`
    rows, err := db.Query(query)
    if err != nil {
      panic(err.Error())
    }
    defer rows.Close()


    var processLists []structs.ProcessList

    for rows.Next() {
      var processList structs.ProcessList

      err := rows.Scan(
        &processList.Id,
        &processList.User,
        &processList.Host,
        &processList.Db,
        &processList.Command,
        &processList.Time,
        &processList.State,
        &processList.Info,
      )
      if err != nil {
        panic(err.Error())
      }

      processLists = append(processLists, processList)
    }

    return processLists
}
