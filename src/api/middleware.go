package api

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

const (
	authBearerSchema string = "Bearer "
)

func UserHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader != "" {
			if !strings.HasPrefix(authHeader, authBearerSchema) {
				http.Error(w, "Forbidden", http.StatusForbidden)
			}
			userToken := authHeader[len(authBearerSchema):]
			token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
				return []byte("AllYourBase"), nil
			})

			if err != nil {
				http.Error(w, "Forbidden", http.StatusForbidden)
			}

			// Is token invalid?
			if !token.Valid {
				http.Error(w, "Forbidden", http.StatusForbidden)
			}

			next.ServeHTTP(w, r)
		}
		http.Error(w, "Forbidden", http.StatusForbidden)

	})
}

func loadData() ([]byte, error) {

	var rdr io.Reader

	if f, err := os.Open(""); err == nil {
		rdr = f
		defer f.Close()
	} else {
		return nil, err
	}

	return ioutil.ReadAll(rdr)
}
