package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type middleware func(http.Handler) http.Handler

type authPostBody struct {
	JWT string `json:"jwt"`
}

// Auth middleware expects a jwt authorization header and verifies it with the
// provided <host>/api/verify
func Auth(host string) middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwt := r.Header.Get("authorization")
			if jwt == "" {
				http.Error(w, `Missing "authorization" header`, http.StatusBadRequest)
				return
			}
			p := &authPostBody{
				JWT: jwt,
			}
			b, err := json.Marshal(p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
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
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		})
	}
}
