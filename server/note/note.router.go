package note

import (
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type RouterData struct {
	Path     string
	Handler  func(pattern string, handler func(http.ResponseWriter, *http.Request))
	Database *mongo.Client
}

func Router(data RouterData) {
	controller := CreateContoller(CreateService(data.Database))

	getPath, getHandler := controller.Get()
	fmt.Println("Загрузка " + data.Path + getPath)
	data.Handler(data.Path+getPath, getHandler)
}
