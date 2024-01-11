package app

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const SHORT_URL_LENGTH = 15

var UrlStorage map[string]string //Storage for shortened URL

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length+2)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[2 : length+2]
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

func reductUrl(url string, shortUrlLength int, urlStorage map[string]string) string {
	shortUrl := randomString(shortUrlLength)
	for {
		_, ok := urlStorage[shortUrl]
		if !ok {
			urlStorage[shortUrl] = url
			return shortUrl
		}
		shortUrl = randomString(shortUrlLength)
	}
}

//Return saved long URL from URL storage
func getUrlFromStorage(id string, urlStorage map[string]string) string {
	url, ok := urlStorage[id]
	if ok {
		//fmt.Printf("DEBUG: Found shorten URL with id '%s' is URL storage.\n", id)
		return url
	} else {
		//fmt.Printf("DEBUG: Shorten URL with id '%s' not found in URL storage.\n", id)
		return ""
	}
}

func PostAndGetHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost: //(i1) Эндпоинт POST / принимает в теле запроса строку URL для сокращения и возвращает ответ с кодом 201 и сокращённым URL в виде текстовой строки в теле.
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		bodyStr := string(body)
		fmt.Printf("DEBUG: POST request body is: '%s'\n", bodyStr)
		w.WriteHeader(http.StatusOK)
		resp := reductUrl(bodyStr, SHORT_URL_LENGTH, UrlStorage)
		fmt.Printf("DEBUG: Shortened URL is: '%s'.\n", resp)
		fmt.Println("DEBUG: UrlStorage:")
		printMap(UrlStorage)
	case http.MethodGet: // Эндпоинт GET /{id} принимает в качестве URL-параметра идентификатор сокращённого URL и возвращает ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.
		urlPath := r.URL.Path
		fmt.Printf("DEBUG: GET method. URL is %s.\n", string(urlPath))
		matched, err := regexp.MatchString(`/[A-Za-z0-9]+$`, urlPath)
		id := trimSlashes(urlPath)
		if matched && (err == nil) && len(id) == SHORT_URL_LENGTH {
			fmt.Printf("DEBUG: Got URL id with lenght %d. URL id = '%s' .\n", SHORT_URL_LENGTH, id)
			longUrl := getUrlFromStorage(id, UrlStorage)
			if longUrl != "" {
				fmt.Printf("DEBUG: Long URL form URL storage with id '%s' is '%s'\n", id, longUrl)
				w.Header().Set("Location", longUrl)
			} else {
				fmt.Printf("DEBUG: Long URL with id '%s' not found in URL storage.", id)
			}
		}
		w.WriteHeader(http.StatusTemporaryRedirect) //code 307
	default:
		fmt.Printf("DEBUG: Only POST and GET request method supported.\n")
		w.WriteHeader(http.StatusBadRequest) //code 400
	}
}
