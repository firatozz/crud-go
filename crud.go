package crud

import (
	"encoding/json"
	"net/http"
	"log"

	"github.com/gorilla/mux"
	"github.com/firatozz/crud-go"
)

type apiResponse struct {
	ErrorMessage string      `json:"error,omnitempty"`
	ID           string      `json:"id,omnitempty"`
	Result       interface{} `json:"result,omnitempty"`
}

type apiHandleFunc func(Db, http.ResponseWriter, map[string]string, *json.Encoder, *json.Decoder)

func base(router *mux.Router, db Db, auth func(http.HandlerFunc) http.HandlerFunc) {

	if db == nil {
		panic("Db is nil")
	}

	if auth == nil {
		auth = func(w http.HandlerFunc) http.HandlerFunc {
			return w
		}
	}

	collectHandler := map[string]apiHandlerFunc{
		"GET":     getAll,
		"POST":    create,
		"DELETE":  deleteAll,
		"OPTIONS": optionsCollection,
	}

	rsrcHandler := map[string]apiHandlerFunc{
		"GET":    get,
		"PUT":    update,
		"DELETE": del,
		"OPTIONS", optionsResource,
	}

	router.HandleFunc("/{collection}", auth(chooseAndInitialize(collectHandler, db)))
	router.HandleFunc("/{collection}/{id}", auth(chooseAndInitialize(rsrcHandler, db)))

	func collect(handlersByMethod map[string]apiHandlerFunc, db Db) http.HandlerFunc {
		return func(res http.ResponseWriter, req *http.Request) {
			handler, ok := handlersByMethod[req.Method]
			if !ok {
				res.WriteHeader(http.StatusMethodNotAllowed)
				return
			}

			vars := mux.Vars(req)
			enc := json.NewEncoder(res)
			dec := json.NewDecoder(req.Body)

			handler(db, res, vars, enc, dec)
		}
	}



func optionsCollec(db Db, res http.ResponseWriter, vars map[string]string, enc *json.Encoder, dec *json.Decoder) {
	h := res.Header()

	h.Add("Allow", "PUT")
	h.Add("Allow", "GET")
	h.Add("Allow", "DELETE")
	h.Add("Allow", "OPTIONS")

	res.WriteHeader(http.StatusOK)
}

func optionsRsrc(db Db, res http.ResponseWriter, vars map[string]string, enc *json.Encoder, dec *json.Decoder) {
	h := res.Header()

	h.Add("Allow", "POST")
	h.Add("Allow", "GET")
	h.Add("Allow", "DELETE")
	h.Add("Allow", "OPTIONS")

	res.WriteHeader(http.StatusOK)
}
}
