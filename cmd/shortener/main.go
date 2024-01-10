package main

import (
	"log"
	"net/http"
	"time"

	"github.com/DmitrySkalnenkov/reduction/internal/app"
)

func main() {
	var hostPortStr string = "localhost:8080" //(i1) Сервер должен быть доступен по адресу: http://localhost:8080.
	s := &http.Server{
		Addr:         hostPortStr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
