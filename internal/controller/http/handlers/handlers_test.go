package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DmitrySkalnenkov/reduction/config"
	"github.com/DmitrySkalnenkov/reduction/internal/controller/memrepo"
	"github.com/DmitrySkalnenkov/reduction/internal/entity"

	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestPostHandler(t *testing.T) {
	sp := config.ServerParameters{
		HostSocketAddrStr: config.DefaultHostURL,
		BaseURLStr:        config.DefaultHostURL,
	}
	config.SetGlobalVariables(&sp)
	entity.URLStorage = new(memrepo.MemRepo)
	entity.URLStorage.InitRepo("")

	entity.URLStorage.SetURLIntoRepo("qwerfadsfd", "https://golang-blog.blogspot.com/2020/01/map-golang.html")
	entity.URLStorage.SetURLIntoRepo("8rewq78rqew", "https://ru.wikipedia.org/wiki/%D0%9A%D0%B8%D1%80%D0%B8%D0%BB%D0%BB%D0%B8%D1%86%D0%B0")
	entity.URLStorage.SetURLIntoRepo("lahfsdafnb4121l", "https://ru.wikipedia.org/wiki/%D0%A3%D0%BC%D0%BB%D0%B0%D1%83%D1%82_(%D0%B4%D0%B8%D0%B0%D0%BA%D1%80%D0%B8%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B9_%D0%B7%D0%BD%D0%B0%D0%BA)")
	entity.URLStorage.SetURLIntoRepo("3123123123123", "https://en.wikipedia.org/wiki/Hungarian_alphabet")
	entity.URLStorage.SetURLIntoRepo("KJFASSFASDJSJ", "https://en.wikipedia.org/wiki/Latin_alphabet")

	type inputStruct struct {
		reqMethod string
		reqURL    string
		reqData   string
		//	urlStorage map[string]string
	}

	type wantStruct struct {
		respCode              int
		respBodyStr           string
		respLocationHeaderStr string
	}

	tests := []struct {
		name  string
		input inputStruct
		want  wantStruct
	}{ //Test table
		{
			name: "Positive test 1. POST request",
			input: inputStruct{
				reqMethod: http.MethodPost,
				reqURL:    "http://127.0.0.1:8080/",
				reqData:   "https://go.dev/tour/moretypes/19",
			},
			want: wantStruct{
				respCode:              http.StatusCreated,
				respBodyStr:           "",
				respLocationHeaderStr: "",
			},
		},
	}
	for _, tt := range tests {
		// запускаем каждый тест
		t.Run(tt.name, func(t *testing.T) {
			reqBytes := []byte(tt.input.reqData)
			req := httptest.NewRequest(tt.input.reqMethod, tt.input.reqURL, bytes.NewReader(reqBytes))
			w := httptest.NewRecorder()
			h := http.HandlerFunc(PostHandler)
			h.ServeHTTP(w, req)
			result := w.Result()
			resultBody, err := io.ReadAll(result.Body)
			if err != nil {
				t.Errorf("TEST_ERROR: %s:", err)
			}
			fmt.Printf("TEST_DEBUG: Response body is '%s'.\n", string(resultBody))
			defer result.Body.Close()

			if result.StatusCode != tt.want.respCode {
				t.Errorf("TEST_ERROR: Expected status code %d, got %d", tt.want.respCode, result.StatusCode)
			} else if len(string(resultBody)) != len(config.DefaultHostURL)+len("/")+config.DefaultShortURLLength {
				t.Errorf("TEST_ERROR: Wrong lenght of result (%d). Should be equal len(DefaultHostURL)+ShortURLLength (%d).\n", len(string(resultBody)), len(config.DefaultHostURL)+config.DefaultShortURLLength)
			}
		})
	}
}
func TestGetHandler(t *testing.T) {
	sp := config.ServerParameters{
		HostSocketAddrStr: config.DefaultHostURL,
		BaseURLStr:        config.DefaultHostURL,
	}
	config.SetGlobalVariables(&sp)
	entity.URLStorage = new(memrepo.MemRepo)
	entity.URLStorage.InitRepo("")

	entity.URLStorage.SetURLIntoRepo("qwerfadsfd", "https://golang-blog.blogspot.com/2020/01/map-golang.html")
	entity.URLStorage.SetURLIntoRepo("8rewq78rqew", "https://ru.wikipedia.org/wiki/%D0%9A%D0%B8%D1%80%D0%B8%D0%BB%D0%BB%D0%B8%D1%86%D0%B0")
	entity.URLStorage.SetURLIntoRepo("lahfsdafnb4121l", "https://ru.wikipedia.org/wiki/%D0%A3%D0%BC%D0%BB%D0%B0%D1%83%D1%82_(%D0%B4%D0%B8%D0%B0%D0%BA%D1%80%D0%B8%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B9_%D0%B7%D0%BD%D0%B0%D0%BA)")
	entity.URLStorage.SetURLIntoRepo("3123123123123", "https://en.wikipedia.org/wiki/Hungarian_alphabet")
	entity.URLStorage.SetURLIntoRepo("KJFASSFASDJSJ", "https://en.wikipedia.org/wiki/Latin_alphabet")

	type inputStruct struct {
		reqMethod string
		reqURL    string
		reqData   string
		//	urlStorage map[string]string
	}

	type wantStruct struct {
		respCode              int
		respBodyStr           string
		respLocationHeaderStr string
	}

	tests := []struct {
		name  string
		input inputStruct
		want  wantStruct
	}{ //Test table
		{
			name: "Positive test 1. GET request (only letters)",
			input: inputStruct{
				reqMethod: http.MethodGet,
				reqURL:    "http://localhost:8080/qwerfadsfd",
				reqData:   "",
			},
			want: wantStruct{
				respCode:              http.StatusTemporaryRedirect,
				respBodyStr:           "",
				respLocationHeaderStr: "https://golang-blog.blogspot.com/2020/01/map-golang.html",
			},
		},
		{
			name: "Positive test 2. GET request (only digits)",
			input: inputStruct{
				reqMethod: http.MethodGet,
				reqURL:    "http://localhost:8080/3123123123123",
				reqData:   "",
			},
			want: wantStruct{
				respCode:              http.StatusTemporaryRedirect,
				respBodyStr:           "",
				respLocationHeaderStr: "https://en.wikipedia.org/wiki/Hungarian_alphabet",
			},
		},
	}
	for _, tt := range tests {
		// запускаем каждый тест
		t.Run(tt.name, func(t *testing.T) {
			reqBytes := []byte(tt.input.reqData)
			var resultLocation *url.URL
			var resultLocationFullURLStr string
			req := httptest.NewRequest(tt.input.reqMethod, tt.input.reqURL, bytes.NewReader(reqBytes))
			w := httptest.NewRecorder()
			h := http.HandlerFunc(GetHandler)
			h.ServeHTTP(w, req)
			result := w.Result()
			resultBody, err := io.ReadAll(result.Body)
			if err != nil {
				t.Errorf("TEST_ERROR: %s:", err)
			}
			defer result.Body.Close()
			fmt.Printf("TEST_DEBUG: Response body is '%s'.\n", string(resultBody))
			resultLocation, err = result.Location()
			if err != nil {
				fmt.Printf("TEST_DEBUG: Cannot get 'Location' from the response. Err - %s.\n", err)
			} else {
				resultLocationFullURLStr = fmt.Sprintf("%v", resultLocation)
				fmt.Printf("TEST_DEBUG: Location header is '%s'.\n", resultLocationFullURLStr)
			}
			if result.StatusCode != tt.want.respCode {
				t.Errorf("TEST_ERROR: Expected status code %d, got %d", tt.want.respCode, result.StatusCode)
			} else if resultLocation != nil {
				if resultLocationFullURLStr != tt.want.respLocationHeaderStr {
					t.Errorf("TEST_ERROR: Location header is '%s' but should be '%s'.\n", string(resultLocationFullURLStr), tt.want.respLocationHeaderStr)
				}
			}
		})
	}
}
func TestPostShortenHandler(t *testing.T) {
	//storage.URLStorage.InitRepo("")
	entity.URLStorage = new(memrepo.MemRepo)
	entity.URLStorage.InitRepo("")

	entity.URLStorage.SetURLIntoRepo("qwerfadsfd", "https://golang-blog.blogspot.com/2020/01/map-golang.html")
	entity.URLStorage.SetURLIntoRepo("8rewq78rqew", "https://ru.wikipedia.org/wiki/%D0%9A%D0%B8%D1%80%D0%B8%D0%BB%D0%BB%D0%B8%D1%86%D0%B0")
	entity.URLStorage.SetURLIntoRepo("lahfsdafnb4121l", "https://ru.wikipedia.org/wiki/%D0%A3%D0%BC%D0%BB%D0%B0%D1%83%D1%82_(%D0%B4%D0%B8%D0%B0%D0%BA%D1%80%D0%B8%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B9_%D0%B7%D0%BD%D0%B0%D0%BA)")
	entity.URLStorage.SetURLIntoRepo("3123123123123", "https://en.wikipedia.org/wiki/Hungarian_alphabet")
	entity.URLStorage.SetURLIntoRepo("KJFASSFASDJSJ", "https://en.wikipedia.org/wiki/Latin_alphabet")

	type inputStruct struct {
		reqMethod      string
		reqURL         string
		reqContentType string
		reqJSONMsg     entity.TxJSONMessage
	}

	type wantStruct struct {
		respCode        int
		respContentType string
		respJSONMsg     entity.TxJSONMessage
	}

	tests := []struct {
		name  string
		input inputStruct
		want  wantStruct
	}{ //Test table
		{
			name: "Positive test 1. POST request",
			input: inputStruct{
				reqMethod:      http.MethodPost,
				reqURL:         "http://127.0.0.1:8080/",
				reqContentType: "application/json",
				reqJSONMsg: entity.TxJSONMessage{
					URL: "http://google.com",
				},
			},
			want: wantStruct{
				respCode:        http.StatusCreated,
				respContentType: "application/json",
			},
		},
		{
			name: "Negative test 1. Wrong header request",
			input: inputStruct{
				reqMethod:      http.MethodPost,
				reqURL:         "http://127.0.0.1:8080/",
				reqContentType: "application/xml",
				reqJSONMsg: entity.TxJSONMessage{
					URL: "http://google.com",
				},
			},
			want: wantStruct{
				respCode:        http.StatusBadRequest,
				respContentType: "application/json",
			},
		},
		{
			name: "Negative test 1. Empty URL in JSON body",
			input: inputStruct{
				reqMethod:      http.MethodPost,
				reqURL:         "http://127.0.0.1:8080/",
				reqContentType: "application/json",
				reqJSONMsg: entity.TxJSONMessage{
					URL: "",
				},
			},
			want: wantStruct{
				respCode:        http.StatusNotFound,
				respContentType: "application/json",
			},
		},
	}
	for _, tt := range tests {
		// запускаем каждый тест
		t.Run(tt.name, func(t *testing.T) {
			reqJSONBody, err := json.Marshal(tt.input.reqJSONMsg)
			if err != nil {
				t.Errorf("TEST_ERROR: '%s'.\n", err)
			}
			req := httptest.NewRequest(tt.input.reqMethod, tt.input.reqURL, bytes.NewReader(reqJSONBody))
			req.Header.Set("Content-Type", tt.input.reqContentType)

			w := httptest.NewRecorder()
			h := http.HandlerFunc(PostShortenHandler)
			h.ServeHTTP(w, req)
			result := w.Result()
			resultBody, err := io.ReadAll(result.Body)
			if err != nil {
				t.Errorf("TEST_ERROR: %s:", err)
			}
			fmt.Printf("TEST_DEBUG: Response body is '%s'.\n", string(resultBody))
			defer result.Body.Close()
			if result.Header.Get("Content-Type") == "application/json" {
				var resMsg entity.TxJSONMessage
				err = json.NewDecoder(w.Body).Decode(&resMsg)
				if err != nil {
					t.Errorf("TEST_ERROR: %s:", err)
				}
			}
			//fmt.Printf("TEST_DEBUG: Response body: URL = '%s\n", resMsg.URL)
			if result.StatusCode != tt.want.respCode {
				t.Errorf("TEST_ERROR: Expected status code %d, got %d", tt.want.respCode, result.StatusCode)
			}
		})
	}
}
