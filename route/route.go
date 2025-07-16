package route

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type Route struct {
	Path     string
	Router   *chi.Mux
	Database *mongo.Client
}

type ReturnResolvedPath struct {
	RouteType string
	RoutePath string
}

type ReturnRoute struct {
	Path    string
	Handler http.HandlerFunc
}

func ResolvePath(route string, path ...string) ReturnResolvedPath {
	r := strings.Split(route, ": ")

	var RouteType = r[0]
	var RoutePath = r[1]

	for _, value := range path {
		RoutePath += value
	}

	RoutePath = strings.ReplaceAll(RoutePath+"/", "//", "/")

	fmt.Println(RouteType, RoutePath)

	return ReturnResolvedPath{RouteType, RoutePath}
}
