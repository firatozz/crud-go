package crud

import (
	"encoding/json"
	"log"
	"net/http"
)

func getAll(db Db, res http.ResponseWriter, vars map[string]string, enc *json.Encoder, dec *json.Decoder) {
	// look for resources
	resources, dbRes := db.GetAll(vars["collection"])

	// write response
	res.WriteHeader(dbRes.StatusCode())
	err := enc.Encode(apiResponse{dbRes.Error(), "", resources})
	if err != nil {
		log.Println(err)
	}
}

func get(db Db, res http.ResponseWriter, vars map[string]string, enc *json.Encoder, dec *json.Decoder) {
	// look for resource
	resource, dbRes := db.Get(vars["collection"], vars["id"])

	// write response
	res.WriteHeader(dbRes.StatusCode())
	err := enc.Encode(apiResponse{dbRes.Error(), "", resource})
	if err != nil {
		log.Println(err)
	}
}
