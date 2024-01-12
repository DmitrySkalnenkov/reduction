package app

import (
	"fmt"
	"io"

	//"math/rand"
	"crypto/rand"
	"encoding/base32"
	"net/http"
	"regexp"
	"strings"
	//	"time"
)

const ShortURLLength = 15
const (
	HostPort string = ":8080"
	HostAddr string = "localhost"
	HostURL  string = "http://" + HostAddr + HostPort + "/"
)

var URLStorage map[string]string //Storage for shortened URL

func randomString(length int) string {
	randomBytes := make([]byte, 32)
	if length > 0 {
		_, err := rand.Read(randomBytes)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Printf("ERROR: Random string length should be > 0. Now it is %d.\n", length)
		return ""
	}
	randFixStr := base32.StdEncoding.EncodeToString(randomBytes)[:length]
	return strings.ToLower(randFixStr)
}
func printMap(mapa map[string]string) {
	for k, v := range mapa {
		fmt.Println(k, "value is", v)
	}
	fmt.Println("")
}

func trimSlashes(slashedStr string) string {
	return strings.ReplaceAll(slashedStr, "/", "")
}

func reductURL(url string, shortURLLength int, urlStorage map[string]string) string {
	shortURL := randomString(shortURLLength)
	for {
		_, ok := urlStorage[shortURL]
		if !ok {
			urlStorage[shortURL] = url
			return shortURL
		}
		shortURL = randomString(shortURLLength)
	}
}

//Return saved long URL from URL storage
func getURLFromStorage(id string, urlStorage map[string]string) string {
	url, ok := urlStorage[id]
	if ok {
		fmt.Printf("DEBUG: Found shorten URL with id '%s' is URL storage.\n", id)
		return url
	} else {
		fmt.Printf("DEBUG: Shorten URL with id '%s' not found in URL storage.\n", id)
		return ""
	}
}

func PostAndGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
	fmt.Println("DEBUG: UrlStorage:")
	printMap(URLStorage)
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
		resp := reductURL(bodyStr, ShortURLLength, URLStorage)
		fmt.Printf("DEBUG: Shortened URL is: '%s'.\n", resp)
		_, err = w.Write([]byte(HostURL + resp))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		return
	case http.MethodGet: // Эндпоинт GET /{id} принимает в качестве URL-параметра идентификатор сокращённого URL и возвращает ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.
		urlPath := r.URL.Path
		fmt.Printf("DEBUG: GET method. URL is %s.\n", string(urlPath))
		matched, err := regexp.MatchString(`/[A-Za-z0-9]+$`, urlPath)
		id := trimSlashes(urlPath)
		fmt.Printf("DEBUG: Token of shortened URL is '%s'.\n", id)
		if matched && (err == nil) {
			fmt.Printf("DEBUG: Got URL id with lenght %d. URL id = '%s' .\n", ShortURLLength, id)
			longURL := getURLFromStorage(id, URLStorage)
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
