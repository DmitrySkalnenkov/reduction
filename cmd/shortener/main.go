package main

import (
	"fmt"

	"log"
	"math/rand"
	"net/http"
	"time"
	//	"github.com/DmitrySkalnenkov/reduction/internal/app"
)

//(i1) Сервер должен предоставлять два эндпоинта: POST / и GET /{id}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length+2)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[2 : length+2]
}

func Reduction(url string, shortUrlLength int, urlStorage map[string]string) {
	shortUrl := randomString(shortUrlLength)
	for {
		_, ok := urlStorage[shortUrl]
		if !ok {
			urlStorage[shortUrl] = url
			fmt.Printf("%v", urlStorage)
			return
		}
		shortUrl = randomString(shortUrlLength)
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request) { //(i1) Эндпоинт GET /{id} принимает в качестве URL-параметра идентификатор сокращённого URL и возвращает ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.
	if r.Method == http.MethodGet {

	} else {
		return
	}
}

/*
func PostHandler(w http.ResponseWriter, r *http.Request) { //(i1) Эндпоинт POST / принимает в теле запроса строку URL для сокращения и возвращает ответ с кодом 201 и сокращённым URL в виде текстовой строки в теле.
	if r.Method == http.MethodPost {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}
*/
func main() {
	SHORT_URL_LENGTH := 15
	var urlStorage map[string]string //Storage for shortened URL
	urlStorage = make(map[string]string)
	url := "http://server.com:8080/adfasdfasd/asfdasfas"
	Reduction(url, SHORT_URL_LENGTH, urlStorage)

	var hostPortStr string = "localhost:8080" //(i1) Сервер должен быть доступен по адресу: http://localhost:8080.
	s := &http.Server{
		Addr:         hostPortStr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
