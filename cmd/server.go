package main

import (
	"log"
	"net/http"
	"verification/pkg/api"
)

func main() {
	// Запуск сетевой службы и HTTP-сервера
	// на всех локальных IP-адресах на порту 80.
	err := http.ListenAndServe(":8083", api.New().Router())
	if err != nil {
		log.Fatal(err)
	}
}
