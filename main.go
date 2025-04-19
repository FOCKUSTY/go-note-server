package main

import (
	"fmt"
	"net/http"

	"main/env"
	server "main/server"
)

func main() {
	env.InitEnv()

	PORT, portExists := env.Get("PORT")

	if !portExists {
		fmt.Println("Ошибка: порт не найден")
		return
	}

	fmt.Println("Загрузка роутинга")
	server.Router("/", http.HandleFunc)

	fmt.Println("Прослушивание на порте: " + PORT)
	fmt.Println("http://localhost:" + PORT)

	http.ListenAndServe(":"+PORT, nil)
}
