package crud

import (
	"encoding/json"
	"log"
	"net/http"
)

func update(db Db, res http.ResponseWriter, vars map[string]string, enc *json.Encoder, dec *json.Decoder) {
	var resource map[string]interface{}
	err := dec.Decode(&resource)

	if err != nil {
		log.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		err := enc.Encode(apiResponse{"malformed json", "", nil})
		if err != nil {
			log.Println(err)
		}

		return
	}

	// update resource
	dbRes := db.Update(vars["collection"], vars["id"], resource)

	// write response
	res.WriteHeader(dbRes.StatusCode())
	err = enc.Encode(apiResponse{dbRes.Error(), "", nil})
	if err != nil {
		log.Println(err)
	}
}
