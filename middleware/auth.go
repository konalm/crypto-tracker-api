package middleware

import (
  "net/http"
  "stelita-api/authentication"
  "github.com/gorilla/context"
)

func Auth(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // get token out of header
    authToken := r.Header.Get("Authorization")
    authTokenResponse := authentication.ValidateToken(authToken)

    if authTokenResponse.Valid {
      // Call the next handler, which can be another middleware in the chain, or the final handler.
      context.Set(r, "userId", authTokenResponse.UserId)
      next.ServeHTTP(w, r)
      return
    }

    w.WriteHeader(406)
    w.Write([]byte("Authorization token is not valid"))
  })
}
