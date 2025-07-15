package main

import (
	"main/env"

	"fmt"
	"net/http"

	"main/database"
	server "main/server"
)

func main() {
	PORT := env.Get("PORT")
	client := database.CreateClient(env.Get("MONGO_URL"))

	fmt.Println("Загрузка роутинга")
	server.Router(server.RouterData{
		Path:     "",
		Handler:  http.HandleFunc,
		Database: client.Client,
	})

	fmt.Println("Прослушивание на порте: " + PORT)
	fmt.Println("http://localhost:" + PORT)

	http.ListenAndServe(":"+PORT, nil)
}
