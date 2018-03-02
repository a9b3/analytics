package v1

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HelloHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("hi")
}
