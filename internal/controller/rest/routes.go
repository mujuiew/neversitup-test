package rest

import (
	"github.com/go-chi/chi"
	"github.com/neversitup-test/internal/controller/rest/shufflings"
)

func Load() {
	var router chi.Router

	router.Post("/shufflings", shufflings.Handle)
}
