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

	fmt.Fprint(w, string(b))
}
