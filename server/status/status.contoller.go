package status

import (
	"encoding/json"
	"main/route"
	"net/http"
)

type Controller struct {
	Path    string
	Routes  *Routes
	Service *Service
}

func (contoller Controller) Get() (string, http.HandlerFunc) {
	return route.AddPrefixToRoute(contoller.Routes.GET, contoller.Path).RoutePath, func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		response := contoller.Service.Get()

		json.NewEncoder(writer).Encode(response)
	}
}

func (contoller Controller) GetDatabase() (string, http.HandlerFunc) {
	return route.AddPrefixToRoute(contoller.Routes.GET_DATABASE, contoller.Path).RoutePath, func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		response := contoller.Service.GetDatabase()

		json.NewEncoder(writer).Encode(response)
	}
}

func CreateContoller(service *Service, Path string) *Controller {
	return &Controller{Routes: GetRoutes(), Service: service, Path: Path}
}
