package http

import (
	"github.com/DmitrySkalnenkov/reduction/internal/controllers/http/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() (m *chi.Mux) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Compress(5))
	r.Post("/", handlers.PostHandler)                   //(i1) Эндпоинт POST / принимает в теле запроса строку URL для сокращения и возвращает ответ с кодом 201 и сокращённым URL в виде текстовой строки в теле.
	r.Get("/{id}", handlers.GetHandler)                 //(i1) Эндпоинт GET /{id} принимает в качестве URL-параметра идентификатор сокращённого URL и возвращает ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.
	r.Post("/api/shorten", handlers.PostShortenHandler) //(i4) Добавьте в сервер новый эндпоинт POST /api/shorten, принимающий в теле запроса JSON-объект {"url":"<some_url>"} и возвращающий в ответ объект {"result":"<shorten_url>"}.
	r.NotFound(r.NotFoundHandler())
	r.MethodNotAllowed(r.MethodNotAllowedHandler())
	return r
}
