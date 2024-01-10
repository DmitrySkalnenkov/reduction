package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
	//	"github.com/DmitrySkalnenkov/reduction/internal/app"
)

//(i1)Сервер должен предоставлять два эндпоинта: POST / и GET /{id}

func GetHandler(w http.ResponseWriter, r *http.Request) { //(i1) Эндпоинт GET /{id} принимает в качестве URL-параметра идентификатор сокращённого URL и возвращает ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.
	if r.Method == http.MethodGet {

	} else {
		return
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) { //(i1) Эндпоинт POST / принимает в теле запроса строку URL для сокращения и возвращает ответ с кодом 201 и сокращённым URL в виде текстовой строки в теле.
	if r.Method == http.MethodPost {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resp, err := json.Marshal(b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}

func main() {
	var hostPortStr string = "localhost:8080" //(i1) Сервер должен быть доступен по адресу: http://localhost:8080.
	s := &http.Server{
		Addr:         hostPortStr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
	http.HandleFunc("/", PostHandler)

	log.Fatal(s.ListenAndServe())
}
