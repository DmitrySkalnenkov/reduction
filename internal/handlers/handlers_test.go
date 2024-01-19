package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/DmitrySkalnenkov/reduction/internal/app"
	"github.com/DmitrySkalnenkov/reduction/internal/storage"
)

func TestPostAndGetHandler(t *testing.T) {
	storage.URLStorage.Init()
	storage.URLStorage.SetURLIntoStorage("qwerfadsfd", "https://golang-blog.blogspot.com/2020/01/map-golang.html")
	storage.URLStorage.SetURLIntoStorage("8rewq78rqew", "https://ru.wikipedia.org/wiki/%D0%9A%D0%B8%D1%80%D0%B8%D0%BB%D0%BB%D0%B8%D1%86%D0%B0")
	storage.URLStorage.SetURLIntoStorage("lahfsdafnb4121l", "https://ru.wikipedia.org/wiki/%D0%A3%D0%BC%D0%BB%D0%B0%D1%83%D1%82_(%D0%B4%D0%B8%D0%B0%D0%BA%D1%80%D0%B8%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B9_%D0%B7%D0%BD%D0%B0%D0%BA)")
	storage.URLStorage.SetURLIntoStorage("3123123123123", "https://en.wikipedia.org/wiki/Hungarian_alphabet")
	storage.URLStorage.SetURLIntoStorage("KJFASSFASDJSJ", "https://en.wikipedia.org/wiki/Latin_alphabet")

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
		{
			name: "Positive test 2. GET request (only letters)",
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
			name: "Positive test 3. GET request (only digits)",
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
			h := http.HandlerFunc(PostAndGetHandler)
			h.ServeHTTP(w, req)
			result := w.Result()
			resultBody, err := io.ReadAll(result.Body)
			if err != nil {
				t.Errorf("TEST_ERROR: %s:", err)
			}
			fmt.Printf("TEST_DEBUG: Response body is '%s'.\n", string(resultBody))
			defer result.Body.Close()
			if tt.input.reqMethod == http.MethodGet {
				resultLocation, err = result.Location()
				if err != nil {
					fmt.Printf("TEST_DEBUG: Cannot get 'Location' from the response. Err - %s.\n", err)
				} else {
					resultLocationFullURLStr = fmt.Sprintf("%v", resultLocation)
					fmt.Printf("TEST_DEBUG: Location header is '%s'.\n", resultLocationFullURLStr)
				}
			}
			if result.StatusCode != tt.want.respCode {
				t.Errorf("TEST_ERROR: Expected status code %d, got %d", tt.want.respCode, result.StatusCode)
			} else if tt.input.reqMethod == http.MethodPost && len(string(resultBody)) != len(app.HostURL)+app.ShortURLLength {
				t.Errorf("TEST_ERROR: Wrong lenght of result (%d). Should be equal len(HostURL)+ShortURLLength (%d).\n", len(string(resultBody)), len(app.HostURL)+app.ShortURLLength)
			} else if tt.input.reqMethod == http.MethodGet && resultLocation != nil {
				if resultLocationFullURLStr != tt.want.respLocationHeaderStr {
					t.Errorf("TEST_ERROR: Location header is '%s' but should be '%s'.\n", string(resultLocationFullURLStr), tt.want.respLocationHeaderStr)
				}
			}
		})
	}
}

func TestPostHandler(t *testing.T) {
	storage.URLStorage.Init()
	storage.URLStorage.SetURLIntoStorage("qwerfadsfd", "https://golang-blog.blogspot.com/2020/01/map-golang.html")
	storage.URLStorage.SetURLIntoStorage("8rewq78rqew", "https://ru.wikipedia.org/wiki/%D0%9A%D0%B8%D1%80%D0%B8%D0%BB%D0%BB%D0%B8%D1%86%D0%B0")
	storage.URLStorage.SetURLIntoStorage("lahfsdafnb4121l", "https://ru.wikipedia.org/wiki/%D0%A3%D0%BC%D0%BB%D0%B0%D1%83%D1%82_(%D0%B4%D0%B8%D0%B0%D0%BA%D1%80%D0%B8%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B9_%D0%B7%D0%BD%D0%B0%D0%BA)")
	storage.URLStorage.SetURLIntoStorage("3123123123123", "https://en.wikipedia.org/wiki/Hungarian_alphabet")
	storage.URLStorage.SetURLIntoStorage("KJFASSFASDJSJ", "https://en.wikipedia.org/wiki/Latin_alphabet")

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
			} else if len(string(resultBody)) != len(app.HostURL)+app.ShortURLLength {
				t.Errorf("TEST_ERROR: Wrong lenght of result (%d). Should be equal len(HostURL)+ShortURLLength (%d).\n", len(string(resultBody)), len(app.HostURL)+app.ShortURLLength)
			}
		})
	}
}
func TestGetHandler(t *testing.T) {
	storage.URLStorage.Init()
	storage.URLStorage.SetURLIntoStorage("qwerfadsfd", "https://golang-blog.blogspot.com/2020/01/map-golang.html")
	storage.URLStorage.SetURLIntoStorage("8rewq78rqew", "https://ru.wikipedia.org/wiki/%D0%9A%D0%B8%D1%80%D0%B8%D0%BB%D0%BB%D0%B8%D1%86%D0%B0")
	storage.URLStorage.SetURLIntoStorage("lahfsdafnb4121l", "https://ru.wikipedia.org/wiki/%D0%A3%D0%BC%D0%BB%D0%B0%D1%83%D1%82_(%D0%B4%D0%B8%D0%B0%D0%BA%D1%80%D0%B8%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B9_%D0%B7%D0%BD%D0%B0%D0%BA)")
	storage.URLStorage.SetURLIntoStorage("3123123123123", "https://en.wikipedia.org/wiki/Hungarian_alphabet")
	storage.URLStorage.SetURLIntoStorage("KJFASSFASDJSJ", "https://en.wikipedia.org/wiki/Latin_alphabet")

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
			fmt.Printf("TEST_DEBUG: Response body is '%s'.\n", string(resultBody))
			defer result.Body.Close()
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
