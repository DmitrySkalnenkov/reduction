package main

import (
	"log"
	"net/http"
	"time"

	"github.com/DmitrySkalnenkov/reduction/internal/app"
)

//(i1) Сервер должен предоставлять два эндпоинта: POST / и GET /{id}

func main() {

	app.UrlStorage = make(map[string]string)

	var hostPortStr string = "localhost:8080" //(i1) Сервер должен быть доступен по адресу: http://localhost:8080.
	s := &http.Server{
		Addr:         hostPortStr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
	http.HandleFunc("/", app.PostAndGetHandler)
	
	log.Fatal(s.ListenAndServe())
}
