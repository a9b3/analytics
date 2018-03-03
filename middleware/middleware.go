package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Adapter func(http.Handler) http.Handler

type authPostBody struct {
	JWT string `json:"jwt"`
}

func Auth(host string) Adapter {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwt := r.Header.Get("authorization")
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
