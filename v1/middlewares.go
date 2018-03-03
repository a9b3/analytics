package v1

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Adapter func(httprouter.Handle) httprouter.Handle

type authPostBody struct {
	JWT string `json:"jwt"`
}

func AuthMiddleware(host string) Adapter {
	return func(h httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			jwt := r.Header.Get("authorization")
			p := &authPostBody{
				JWT: jwt,
			}
			b, err := json.Marshal(p)
			if err != nil {
				panic(err)
			}

			resp, err := http.Post("http://"+host+"/api/verify", "application/json", bytes.NewBuffer(b))
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			content, _ := ioutil.ReadAll(resp.Body)
			if string(content) == "true" {
				h(w, r, ps)
			} else {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			}
		}
	}
}
