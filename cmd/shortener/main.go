package main

import (
	"log"
	"net/http"

	"github.com/DmitrySkalnenkov/reduction/internal/app"
)

//(i1) Сервер должен предоставлять два эндпоинта: POST / и GET /{id}

func main() {

	app.URLStorage = make(map[string]string)
	var hostPortStr = app.HostAddr + app.HostPort //(i1) Сервер должен быть доступен по адресу: http://localhost:8080.

	s := &http.Server{
		Addr: hostPortStr,
	}
	http.HandleFunc("/", app.PostAndGetHandler)

	log.Fatal(s.ListenAndServe())
}
