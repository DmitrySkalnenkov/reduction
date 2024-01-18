package main

import (
	"log"
	"net/http"

	"github.com/DmitrySkalnenkov/reduction/internal/app"
	"github.com/DmitrySkalnenkov/reduction/internal/handlers"
	"github.com/DmitrySkalnenkov/reduction/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//(i1) Сервер должен предоставлять два эндпоинта: POST / и GET /{id}
func main() {

	storage.URLStorage = make(map[string]string)
	var hostPortStr = app.HostAddr + app.HostPort //(i1) Сервер должен быть доступен по адресу: http://localhost:8080.
	//http.HandleFunc("/", app.PostAndGetHandler) // (i3) Вы написали приложение с помощью стандартной библиотеки net/http. Используя любой пакет(роутер или фреймворк), совместимый с net/http, перепишите ваш код.
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/", handlers.PostHandler)
	r.Get("/{id}", handlers.GetHandler)
	r.NotFound(r.NotFoundHandler())
	r.MethodNotAllowed(r.MethodNotAllowedHandler())

	s := &http.Server{
		Addr: hostPortStr,
	}
	s.Handler = r
	log.Fatal(s.ListenAndServe())
}
