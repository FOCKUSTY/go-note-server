package hello

import (
	"fmt"
	"net/http"
)

var routes = CreateRoutes()
var controller = CreateController()

func Router(path string, handler func(pattern string, handler func(http.ResponseWriter, *http.Request))) {
	fmt.Println("Загрузка " + path + routes.GET)

	handler(path+routes.GET, controller.Get)
}
