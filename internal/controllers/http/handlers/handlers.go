package handlers

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/DmitrySkalnenkov/reduction/config"
	"github.com/DmitrySkalnenkov/reduction/internal/controllers/http/cookies"
	"github.com/DmitrySkalnenkov/reduction/internal/models"
	"github.com/DmitrySkalnenkov/reduction/internal/services"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
)

type gzWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w gzWriter) Write(b []byte) (int, error) {
	// w.Writer будет отвечать за gzip-сжатие, поэтому пишем в него
	return w.Writer.Write(b)
}

func GetRequestBody(w http.ResponseWriter, r *http.Request) (bodyStr string, respStatus int) {
	models.URLStorage.PrintRepo() //for DEBUG
	var curJSONMsg models.TxJSONMessage
	var reader io.Reader
	switch r.Header.Get("Content-Type") {
	case "application/json":
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&curJSONMsg)
		if err != nil {
			log.Printf("ERROR: JSON decoding occured, %s.\n", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return "", http.StatusBadRequest
		} else if curJSONMsg.URL != "" {
			bodyStr = curJSONMsg.URL
			fmt.Printf("DEBUG: JSON body: URL = '%s'.\n", bodyStr)
			respStatus = http.StatusCreated
		} else {
			bodyStr = curJSONMsg.URL
			respStatus = http.StatusNotFound
		}
	case "", "text/plain", "text/plain; charset=utf-8", "text/html; charset=utf-8":
		if r.Header.Get("Content-Encoding") == "gzip" {
			gz, err := gzip.NewReader(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				fmt.Printf("ERROR: %v", err)
				return "", http.StatusInternalServerError
			}
			reader = gz
			defer gz.Close()
		} else {
			reader = r.Body
		}
		body, err := io.ReadAll(reader)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return "", http.StatusBadRequest
		}
		bodyStr = string(body)
		fmt.Printf("DEBUG: POST request body is: '%s'\n", bodyStr)
		respStatus = http.StatusCreated
	default:
		bodyStr = ""
		respStatus = http.StatusBadRequest
	}
	return bodyStr, respStatus
}

// GET handler
func GetHandler(w http.ResponseWriter, r *http.Request) {
	models.URLStorage.PrintRepo() //For Debug
	id := chi.URLParam(r, "id")
	fmt.Printf("DEBUG: Token of shortened URL is '%s'.\n", id)
	longURL, ok := models.URLStorage.GetURLFromRepo(id)
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

// POST handler
func PostHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	curBodyStr, curRespStatus := GetRequestBody(w, r)
	curUserID := cookies.GetAuthUserID(w, r)
	curURLUser := models.URLUser{URL: curBodyStr, UserID: curUserID}
	token := services.ReduceURL(curURLUser, config.DefaultShortURLLength, models.URLStorage)
	fmt.Printf("DEBUG: Shortened URL for UserID = %d is: '%s'.\n", curUserID, token)
	models.URLStorage.PrintRepo() //for DEBUG
	w.WriteHeader(curRespStatus)
	_, err = w.Write([]byte(models.BaseURLStr + "/" + token))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func PostShortenHandler(w http.ResponseWriter, r *http.Request) {
	var respJSONMsg models.RxJSONMessage
	curBodyStr, curRespStatus := GetRequestBody(w, r)
	curUserID := cookies.GetAuthUserID(w, r)
	curURLUser := models.URLUser{URL: curBodyStr, UserID: curUserID}
	token := services.ReduceURL(curURLUser, config.DefaultShortURLLength, models.URLStorage)
	shortenURL := models.BaseURLStr + "/" + token
	fmt.Printf("DEBUG: Shortened URL for UserID = %d is: '%s'.\n", curUserID, token)
	respJSONMsg.Result = shortenURL
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(curRespStatus)
	err := json.NewEncoder(w).Encode(respJSONMsg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

/*
func (h plainHandler) GetRequestBody(w http.ResponseWriter, r *http.Request) (bodyStr string, httpStatus int) {
	reader := r.Body
	body, err := io.ReadAll(reader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return "", http.StatusBadRequest
	}
	bodyStr = string(body)
	fmt.Printf("DEBUG: Got plain request body: '%s'\n", bodyStr)
	return bodyStr, http.StatusCreated
}

func (h plainHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	curBodyStr, curRespStatus := h.GetRequestBody(w, r)
	curUserID := cookies.GetAuthUserID(w, r)
	curURLUser := models.URLUser{URL: curBodyStr, UserID: curUserID}
	token := services.ReduceURL(curURLUser, config.DefaultShortURLLength, models.URLStorage)
	fmt.Printf("DEBUG: Shortened URL for UserID = %d is: '%s'.\n", curUserID, token)
	w.WriteHeader(curRespStatus)
	_, err := w.Write([]byte(models.BaseURLStr + "/" + token))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type PlainHandler http.HandlerFunc{

}



func NewPlanHandler() http.HandlerFunc{
	var ph http.HandlerFunc
	return ph
}

func (h jsonHandler) GetRequestBody(w http.ResponseWriter, r *http.Request) (bodyStr string, respStatus int) {
	var body models.TxJSONMessage
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		log.Printf("ERROR: JSON decoding occured, %s.\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return "", http.StatusBadRequest
	} else if body.URL != "" {
		bodyStr = body.URL
		fmt.Printf("DEBUG: JSON body: URL = '%s'.\n", bodyStr)
		respStatus = http.StatusCreated
	} else {
		bodyStr = ""
		respStatus = http.StatusNotFound
	}
	return bodyStr, respStatus
}

func (h jsonHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	curBodyStr, curRespStatus := h.GetRequestBody(w, r)
	curUserID := cookies.GetAuthUserID(w, r)
	curURLUser := models.URLUser{URL: curBodyStr, UserID: curUserID}
	token := services.ReduceURL(curURLUser, config.DefaultShortURLLength, models.URLStorage)
	fmt.Printf("DEBUG: Shortened URL for UserID = %d is: '%s'.\n", curUserID, token)
	w.WriteHeader(curRespStatus)
	_, err := w.Write([]byte(models.BaseURLStr + "/" + token))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
*/
