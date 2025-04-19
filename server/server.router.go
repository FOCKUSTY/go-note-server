package server

import (
	"fmt"
	"net/http"

	hello "main/server/hello"
)

type Path struct {
	path    string
	handler func(path string, handler func(pattern string, handler func(http.ResponseWriter, *http.Request)))
}

func Router(path string, handler func(pattern string, handler func(http.ResponseWriter, *http.Request))) {
	paths := []Path{
		{path: path + "hello", handler: hello.Router},
	}

	for _, path := range paths {
		fmt.Println("Загрузка " + path.path)
		path.handler(path.path, handler)
	}
}
