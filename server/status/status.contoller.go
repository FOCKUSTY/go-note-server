package status

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
	Routes  *Routes
	Service *Service
}

func (contoller Controller) Get() (string, func(writer http.ResponseWriter, request *http.Request)) {
	return contoller.Routes.GET, func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		response := contoller.Service.Get()

		json.NewEncoder(writer).Encode(response)
	}
}

func (contoller Controller) GetDatabase() (string, func(writer http.ResponseWriter, request *http.Request)) {
	return contoller.Routes.GET_DATABASE, func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		response := contoller.Service.GetDatabase()

		json.NewEncoder(writer).Encode(response)
	}
}

func CreateContoller(service *Service) *Controller {
	return &Controller{Routes: GetRoutes(), Service: service}
}
