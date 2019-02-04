package middleware

import (
  "fmt"
  "net/http"
  "stelita-api/authentication"
)

func Auth(next http.Handler) http.Handler {
  fmt.Println("Auth Middleware")

  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // get token out of header
    authToken := r.Header.Get("Authorization")
    validAuthToken := authentication.ValidateToken(authToken)

    if validAuthToken {
      // Call the next handler, which can be another middleware in the chain, or the final handler.
      next.ServeHTTP(w, r)
      return
    }

    w.WriteHeader(406)
    w.Write([]byte("Authorization token is not valid"))
  })
}
