package note

import (
	"main/route"

	"github.com/go-chi/chi/v5"
)

func Router(data *route.Route) *chi.Mux {
	router := chi.NewRouter()

	controller := CreateContoller(CreateService(data.Database), router, data.Path)

	router.Get(controller.Get())

	return router
}
