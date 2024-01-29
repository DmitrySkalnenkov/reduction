package http

import (
	"github.com/DmitrySkalnenkov/reduction/internal/controller/http/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() (m *chi.Mux) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Compress(5))
	r.Post("/", handlers.PostHandler)
	r.Get("/{id}", handlers.GetHandler)
	r.Post("/api/shorten", handlers.PostShortenHandler) //(i4) Добавьте в сервер новый эндпоинт POST /api/shorten, принимающий в теле запроса JSON-объект {"url":"<some_url>"} и возвращающий в ответ объект {"result":"<shorten_url>"}.
	r.NotFound(r.NotFoundHandler())
	r.MethodNotAllowed(r.MethodNotAllowedHandler())
	return r
}
