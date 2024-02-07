package main

import (
	"github.com/DmitrySkalnenkov/reduction/internal/app"
	"github.com/DmitrySkalnenkov/reduction/internal/handlers"
	"github.com/DmitrySkalnenkov/reduction/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

// (i1) Сервер должен предоставлять два эндпоинта: POST / и GET /{id}
func main() {
	var sp app.ServerParameters
	sp.GetEnvs()
	sp.GetFlags()
	sp.CheckParamPriority()
	storage.URLStorage = storage.URLStorageInit(app.RepoFilePathStr)
	//storage.URLStorage.InitRepo()
	//http.HandleFunc("/", app.PostAndGetHandler) // (i3) Вы написали приложение с помощью стандартной библиотеки net/http. Используя любой пакет(роутер или фреймворк), совместимый с net/http, перепишите ваш код.
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Compress(5))
	r.Post("/", handlers.PostHandler)
	r.Get("/{id}", handlers.GetHandler)
	r.Post("/api/shorten", handlers.PostShortenHandler) //(i4) Добавьте в сервер новый эндпоинт POST /api/shorten, принимающий в теле запроса JSON-объект {"url":"<some_url>"} и возвращающий в ответ объект {"result":"<shorten_url>"}.
	r.NotFound(r.NotFoundHandler())
	r.MethodNotAllowed(r.MethodNotAllowedHandler())

	//storage.URLStorage.RestoreRepoFromJSONFile(app.DefaultRepoFilePath)
	s := &http.Server{
		Addr: app.HostSocketAddrStr,
	}
	s.Handler = r
	log.Fatal(s.ListenAndServe())
}
