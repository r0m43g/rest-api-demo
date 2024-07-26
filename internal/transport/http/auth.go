package http

import (
	"fmt"
	"net/http"
	"strings"
	log "github.com/sirupsen/logrus"

	jwt "gopkg.in/golang-jwt/jwt.v3"
)

func JWTAuthMiddleware(
  originalHandler func(w http.ResponseWriter, r *http.Request),
) func(w http.ResponseWriter, r *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    token := r.Header.Get("Authorization")

    if token == "" {
      log.Warnln("missing Authorization header")
      w.WriteHeader(http.StatusUnauthorized)
      return
    }

    parts := strings.Split(token, " ")
    if len(parts) != 2 || parts[0] != "Bearer" {
      log.Warn("invalid Authorization header")
      w.WriteHeader(http.StatusUnauthorized)
      return
    }
    if !validateToken(parts[1]) {
      log.Warn("unauthorized access")
      w.WriteHeader(http.StatusUnauthorized)
      return
    }

    originalHandler(w, r)
  }
}

func validateToken(accessToken string) bool {
  var myKey = []byte("middleware")
  token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
    }
    return myKey, nil
  })
  return err == nil && token.Valid
}
