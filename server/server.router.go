package server

import (
	"main/server/note"
	"main/server/status"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type RouterData struct {
	Path     string
	Handler  func(pattern string, handler func(http.ResponseWriter, *http.Request))
	Database *mongo.Client
}

func Router(data RouterData) {
	status.Router(status.RouterData{
		Path:     data.Path + "/status",
		Handler:  data.Handler,
		Database: data.Database,
	})

	note.Router(note.RouterData{
		Path:     data.Path + "/note",
		Handler:  data.Handler,
		Database: data.Database,
	})
}
