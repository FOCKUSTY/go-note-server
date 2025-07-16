package server

import (
	"main/route"
	"main/server/note"

	"github.com/go-chi/chi/v5"
)

func Router(data *route.Route) *chi.Mux {
	router := chi.NewRouter()

	router.Mount(data.Path, note.Router(&route.Route{
		Path:     data.Path + "note",
		Database: data.Database,
		Router:   router,
	}))

	return router
}
