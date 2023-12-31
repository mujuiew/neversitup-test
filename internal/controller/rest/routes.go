package rest

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/neversitup-test/internal/controller/rest/odd"
	"github.com/neversitup-test/internal/controller/rest/shufflings"
	"github.com/neversitup-test/internal/controller/rest/smiley"
)

func Load() {

	router := mux.NewRouter()

	router.HandleFunc("/shufflings", shufflings.Handle).Methods("POST")
	router.HandleFunc("/odd-number", odd.Handle).Methods("POST")
	router.HandleFunc("/smiley", smiley.Handle).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
