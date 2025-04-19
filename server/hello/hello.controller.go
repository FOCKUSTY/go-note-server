package hello

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
	Service *Service
}

func CreateController() *Controller {
	return &Controller{Service: CreateService()}
}

func (c Controller) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := c.Service.Get()

	json.NewEncoder(w).Encode(response)
}
