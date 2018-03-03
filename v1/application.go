package v1

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
)

type appHandlers struct {
	Get   httprouter.Handle
	Post  httprouter.Handle
	Patch httprouter.Handle
	Track httprouter.Handle
}

// CreateAppHandlers returns appHandlers
func CreateAppHandlers(db *mgo.Database) appHandlers {
	return appHandlers{
		Get: func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
			fmt.Fprint(w, "hello")
		},
		Post: func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
			fmt.Fprint(w, "hello")
		},
		Patch: func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
			fmt.Fprint(w, "hello")
		},
		Track: func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
			fmt.Fprint(w, "hello")
		},
	}
}
