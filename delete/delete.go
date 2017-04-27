package crud

import (
	"encoding/json"
	"log"
	"net/http"
)

func deleteAll(db Db, res http.ResponseWriter, vars map[string]string, enc *json.Encoder, dec *json.Decoder) {
	// look for resources
	dbRes := db.DeleteAll(vars["collection"])

	// write response
	res.WriteHeader(dbRes.StatusCode())
	err := enc.Encode(apiResponse{dbRes.Error(), "", nil})
	if err != nil {
		log.Println(err)
	}
}

// delete() is a built-in function, thus del() is used here
func del(db Db, res http.ResponseWriter, vars map[string]string, enc *json.Encoder, dec *json.Decoder) {
	// delete resource
	dbRes := db.Delete(vars["collection"], vars["id"])

	// write response
	res.WriteHeader(dbRes.StatusCode())
	err := enc.Encode(apiResponse{dbRes.Error(), "", nil})
	if err != nil {
		log.Println(err)
	}
}
