package authentication

import (
  "fmt"
  "net/http"
)

/**
 *
 */
func AuthCheck(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Auth Check")

  authToken := r.Header.Get("Authorization")
  authTokenResponse := ValidateToken(authToken)

  if !authTokenResponse.Valid {
    w.WriteHeader(406)
    w.Write([]byte("Not Authorized"))
    return
  }

  w.WriteHeader(200)
  w.Write([]byte("Authorized"))
}
