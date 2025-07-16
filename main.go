package main

import (
	"main/env"
	"main/route"

	"fmt"
	"net/http"

	"main/database"
	server "main/server"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	PORT := env.Get("PORT")
	client := database.CreateClient(env.Get("MONGO_URL"))
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	fmt.Println("Загрузка роутинга")
	router.Mount("/", server.Router(&route.Route{
		Path:     "/",
		Database: client.Client,
		Router:   router,
	}))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	fmt.Println("Прослушивание на порте: " + PORT)
	fmt.Println("http://localhost:" + PORT)

	http.ListenAndServe(":"+PORT, router)
}
