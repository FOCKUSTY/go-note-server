package server

import (
	"main/route"
	"main/server/note"
	"main/server/status"

	"github.com/go-chi/chi/v5"
)

func Router(data *route.Route) *chi.Mux {
	router := chi.NewRouter()

	note.Router(&route.Route{
		Path:     data.Path + "note",
		Database: data.Database,
		Router:   router,
	})

	status.Router(&route.Route{
		Path:     data.Path + "status",
		Database: data.Database,
		Router:   router,
	})

	return router
}
