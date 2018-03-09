package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	jwtGo "github.com/dgrijalva/jwt-go"
)

type authPostBody struct {
	JWT string `json:"jwt"`
}

// Auth middleware expects a jwt authorization header and verifies it with the
// provided <host>/api/verify
func Auth(host, tokenSecret string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var ctx context.Context

			jwt := r.Header.Get("authorization")
			if jwt == "" {
				http.Error(w, `Missing "authorization" header`, http.StatusBadRequest)
				return
			}
			p := &authPostBody{JWT: jwt}
			b, err := json.Marshal(p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			token, err := jwtGo.Parse(jwt, func(token *jwtGo.Token) (interface{}, error) {
				return []byte(tokenSecret), nil
			})
			if err != nil || !token.Valid {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			claims, ok := token.Claims.(jwtGo.MapClaims)
			if ok && claims["id"] != "" {
				ctx = context.WithValue(r.Context(), "userId", claims["id"])
			} else {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			resp, err := http.Post("http://"+host+"/api/verify", "application/json", bytes.NewBuffer(b))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer resp.Body.Close()

			content, _ := ioutil.ReadAll(resp.Body)

			if string(content) == "true" {
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		})
	}
}
