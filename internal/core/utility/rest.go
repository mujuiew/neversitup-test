package utility

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func WriteBody(w http.ResponseWriter, obj interface{}) {
	b, err := json.Marshal(obj)
	if err != nil {
		log.Print("ERROR: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprint(w, string(b))
}

func WriteError(w http.ResponseWriter, obj interface{}) {
	b, err := json.Marshal(obj)
	if err != nil {
		log.Print("ERROR: ", err)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, string(b))
}
