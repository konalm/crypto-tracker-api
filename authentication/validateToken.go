package authentication

import (
  "fmt"
  "os"
  "github.com/dgrijalva/jwt-go"
)

type TokenResponse struct {
  Valid bool
  UserId int
}

/**
 *
 */
func ValidateToken(tokenString string) TokenResponse {
  mySigningKey := []byte(os.Getenv("TOKEN_SECRET"))

  // Parse takes the token string and a function for looking up the key. The latter is especially
  // useful if you use multiple keys for your application.  The standard is to use 'kid' in the
  // head of the token to identify which key to use, but the parsed token (head and claims) is provided
  // to the callback, providing flexibility.
  token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    // Don't forget to validate the alg is what you expect:
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
    }

    return mySigningKey, nil
  })

  if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    claimsUserId := claims["UserId"]
    var userIdFloat float64 = claimsUserId.(float64)
    var userId int = int(userIdFloat)

    return TokenResponse {true, userId}
  }

  return TokenResponse {false, 0}
}
