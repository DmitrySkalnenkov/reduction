package handlers

import (
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/DmitrySkalnenkov/reduction/internal/app"
	"github.com/DmitrySkalnenkov/reduction/internal/storage"
	"github.com/go-chi/chi/v5"
)

//POST handler
func PostHandler(w http.ResponseWriter, r *http.Request) {
	storage.PrintMap(storage.URLStorage) //for DEBUG
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bodyStr := string(body)
	fmt.Printf("DEBUG: POST request body is: '%s'\n", bodyStr)
	w.WriteHeader(http.StatusCreated) //code 201
	resp := app.ReductURL(bodyStr, app.ShortURLLength, storage.URLStorage)
	fmt.Printf("DEBUG: Shortened URL is: '%s'.\n", resp)
	storage.PrintMap(storage.URLStorage) //for DEBUG
	_, err = w.Write([]byte(app.HostURL + resp))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

//GET handler
func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
	fmt.Println("DEBUG: UrlStorage:")
	storage.PrintMap(storage.URLStorage)
	id := chi.URLParam(r, "id")
	fmt.Printf("DEBUG: Token of shortened URL is '%s'.\n", id)
	longURL := storage.GetURLFromStorage(id, storage.URLStorage)
	if longURL != "" {
		fmt.Printf("DEBUG: Long URL form URL storage with id '%s' is '%s'\n", id, longURL)
		w.Header().Set("Location", longURL)
	} else {
		fmt.Printf("DEBUG: Long URL with id '%s' not found in URL storage.\n", id)
	}
	w.WriteHeader(http.StatusTemporaryRedirect) //code 307

}

//Handler for not implemented requests
func NotImplementedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("DEBUG: Only POST and GET request method supported.\n")
	w.WriteHeader(http.StatusBadRequest) //code 400

}

//POST and GET handler (legacy)
func PostAndGetHandler(w http.ResponseWriter, r *http.Request) {
	storage.PrintMap(storage.URLStorage) //for DEBUG
	switch r.Method {
	case http.MethodPost: //(i1) Эндпоинт POST / принимает в теле запроса строку URL для сокращения и возвращает ответ с кодом 201 и сокращённым URL в виде текстовой строки в теле.
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		bodyStr := string(body)
		fmt.Printf("DEBUG: POST request body is: '%s'\n", bodyStr)
		w.WriteHeader(http.StatusCreated) //code 201
		resp := app.ReductURL(bodyStr, app.ShortURLLength, storage.URLStorage)
		storage.PrintMap(storage.URLStorage) //for DEBUG
		fmt.Printf("DEBUG: Shortened URL is: '%s'.\n", resp)
		_, err = w.Write([]byte(app.HostURL + resp))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		return
	case http.MethodGet: // Эндпоинт GET /{id} принимает в качестве URL-параметра идентификатор сокращённого URL и возвращает ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.
		urlPath := r.URL.Path
		fmt.Printf("DEBUG: GET method. URL is %s.\n", string(urlPath))
		matched, err := regexp.MatchString(`/[A-Za-z0-9]+$`, urlPath)
		id := app.TrimSlashes(urlPath)
		fmt.Printf("DEBUG: Token of shortened URL is '%s'.\n", id)
		if matched && (err == nil) {
			fmt.Printf("DEBUG: Got URL id with lenght %d. URL id = '%s' .\n", app.ShortURLLength, id)
			longURL := storage.GetURLFromStorage(id, storage.URLStorage)
			if longURL != "" {
				fmt.Printf("DEBUG: Long URL form URL storage with id '%s' is '%s'\n", id, longURL)
				w.Header().Set("Location", longURL)
			} else {
				fmt.Printf("DEBUG: Long URL with id '%s' not found in URL storage.", id)
			}
		}
		w.WriteHeader(http.StatusTemporaryRedirect) //code 307
		return
	default:
		fmt.Printf("DEBUG: Only POST and GET request method supported.\n")
		w.WriteHeader(http.StatusBadRequest) //code 400
		return
	}
}
