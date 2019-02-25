package user

import (
  "database/sql"
  "stelita-api/db"
)

type User struct {
  Id int
  Username string
  Password string
}

/**
 *
 */
func GetUserByUsername(username string) User {
  db := db.Conn()
  defer db.Close()

  query := "SELECT id, username, password FROM users WHERE username = ? AND admin = 1";
  stmt := db.QueryRow(query, username)

  var user User
  err := stmt.Scan(&user.Id, &user.Username, &user.Password)
  if err != nil {
    if err == sql.ErrNoRows {
      return User{Username: "", Password: ""}
    }

    panic(err.Error())
  }

  return user
}
