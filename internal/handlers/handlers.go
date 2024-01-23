package handlers

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"

	"github.com/DmitrySkalnenkov/reduction/internal/app"
	"github.com/DmitrySkalnenkov/reduction/internal/storage"
	"github.com/go-chi/chi/v5"
)

type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w gzipWriter) Write(b []byte) (int, error) {
	// w.Writer будет отвечать за gzip-сжатие, поэтому пишем в него
	return w.Writer.Write(b)
}

// POST handler
func PostHandler(w http.ResponseWriter, r *http.Request) {
	storage.URLStorage.PrintMap() //for DEBUG
	var reader io.Reader

	if r.Header.Get("Content-Encoding") == "gzip" {
		gz, err := gzip.NewReader(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Printf("ERROR: %v", err)
			return
		}
		reader = gz
		defer gz.Close()
	} else {
		reader = r.Body
	}

	body, err := io.ReadAll(reader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bodyStr := string(body)
	fmt.Printf("DEBUG: POST request body is: '%s'\n", bodyStr)
	w.WriteHeader(http.StatusCreated) //code 201
	resp := app.ReduceURL(bodyStr, app.ShortURLLength, &storage.URLStorage)
	fmt.Printf("DEBUG: Shortened URL is: '%s'.\n", resp)
	storage.URLStorage.PrintMap() //for DEBUG
	_, err = w.Write([]byte(app.BaseURLStr + "/" + resp))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

// GET handler
func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
	fmt.Println("DEBUG: UrlStorage:")
	storage.URLStorage.PrintMap()
	id := chi.URLParam(r, "id")
	fmt.Printf("DEBUG: Token of shortened URL is '%s'.\n", id)
	longURL, ok := storage.URLStorage.GetURLFromStorage(id)
	if longURL != "" && ok {
		fmt.Printf("DEBUG: Long URL form URL storage with id '%s' is '%s'\n", id, longURL)
		w.Header().Set("Location", longURL)
	} else {
		fmt.Printf("DEBUG: Long URL with id '%s' not found in URL storage.\n", id)
	}
	w.Header().Set("Accept-Encoding", "gzip")
	w.WriteHeader(http.StatusTemporaryRedirect) //code 307

}

// Handler for not implemented requests
func NotImplementedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("DEBUG: Only POST and GET request method supported.\n")
	w.WriteHeader(http.StatusBadRequest) //code 400

}

// POST and GET handler (legacy)
func PostAndGetHandler(w http.ResponseWriter, r *http.Request) {
	storage.URLStorage.PrintMap() //for DEBUG
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
		resp := app.ReduceURL(bodyStr, app.ShortURLLength, &storage.URLStorage)
		storage.URLStorage.PrintMap() //for DEBUG
		fmt.Printf("DEBUG: Shortened URL is: '%s'.\n", resp)
		_, err = w.Write([]byte(app.BaseURLStr + "/" + resp))
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
			longURL, ok := storage.URLStorage.GetURLFromStorage(id)
			if longURL != "" && ok {
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

func PostShortenHandler(w http.ResponseWriter, r *http.Request) {
	storage.URLStorage.PrintMap() //for DEBUG
	var curJSONMsg storage.TxJSONMessage
	var respJSONMsg storage.RxJSONMessage
	if r.Header.Get("Content-Type") == "application/json" && r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&curJSONMsg)
		if err != nil {
			log.Printf("ERROR: JSON decoding occured, %s.\n", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else if curJSONMsg.URL != "" {
			bodyStr := curJSONMsg.URL
			log.Printf("DEBUG: JSON body: URL = '%s'.\n", bodyStr)
			token := app.ReduceURL(bodyStr, app.ShortURLLength, &storage.URLStorage)
			shortenURL := app.BaseURLStr + "/" + token
			storage.URLStorage.PrintMap() //for DEBUG
			fmt.Printf("DEBUG: Shortened URL is: '%s'.\n", shortenURL)
			respJSONMsg.Result = shortenURL
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated) //code 201
			json.NewEncoder(w).Encode(respJSONMsg)
			return
		} else {
			log.Printf("DEBUG: Wrong JSON body value or not found URL in request body.\n")
			w.WriteHeader(http.StatusNotFound) //code 404
			return
		}
	} else {
		log.Printf("DEBUG: Wrong JSON header or body value.\n")
		w.WriteHeader(http.StatusBadRequest) //code 400
		return
	}
}
