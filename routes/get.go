package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"maxves_task/types"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var keyrow, valuerow []string
	var idcolindex int
	jsonmap := make(map[string]string)

	keyrow = types.Records[0]

	for i, col_name := range keyrow {
		if col_name == "id" {
			idcolindex = i
		}
	}

	log.Printf("idcolindex = %s", idcolindex)

	for _, row := range types.Records[1:] {
		if row[idcolindex] == id {
			valuerow = row
			break
		}
	}

	log.Printf("valuerow = %s", valuerow)
	if len(valuerow) == 0 {
		log.Printf("no valuerow with the requested id")
		_, err := w.Write([]byte("no valuerow with the requested id"))
		if err != nil {
			log.Printf("failed to write the entry to response, err: %s", err)
		}
		return
	}

	for i, key := range keyrow {
		jsonmap[key] = valuerow[i]
	}

	log.Printf("jsonmap = %s", jsonmap)

	data, err := json.Marshal(jsonmap)
	if err != nil {
		log.Printf("failed to marshal map into json bytes, err: %s", err)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		log.Printf("failed to write the entry to response, err: %s", err)
		return
	}
	return
}
