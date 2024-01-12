package app

import (
	//	"net/http/httptest"

	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

//(i2) Покройте сервис юнит-тестами. Сконцентрируйтесь на покрытии тестами эндпоинтов, чтобы защитить API сервиса от случайных изменений.
func TestRandomString(t *testing.T) {
	tests := []struct {
		name     string
		inputVal int
		wantVal  int
	}{ //Test table
		{
			name:     "Positive test. Lenght 10",
			inputVal: 10,
			wantVal:  10,
		},
		{
			name:     "Positive test. Lenght 2",
			inputVal: 2,
			wantVal:  2,
		},
		{
			name:     "Positive test. Lenght 0",
			inputVal: 0,
			wantVal:  0,
		},
		{
			name:     "Negative test",
			inputVal: -2,
			wantVal:  0,
		},
	}
	var resultStr string
	for _, tt := range tests {
		// запускаем каждый тест
		t.Run(tt.name, func(t *testing.T) {
			resultStr = randomString(tt.inputVal)
			fmt.Printf("RandomString(%d) result is: '%s'\n", tt.inputVal, resultStr)
			if len(resultStr) != tt.wantVal {
				t.Errorf("TEST_ERROR: input value is %d,  want is %d ", tt.inputVal, tt.wantVal)
			}
		})
	}
}

func TestTrimSlashes(t *testing.T) {
	tests := []struct {
		name     string
		inputVal string
		wantVal  string
	}{ //Test table
		{
			name:     "Positive test. Slash in the begining of string",
			inputVal: "/test_string1",
			wantVal:  "test_string1",
		},
		{
			name:     "Positive test. Slash in the end of string",
			inputVal: "test_string2/",
			wantVal:  "test_string2",
		},
		{
			name:     "Positive test. Slash in the middle of string",
			inputVal: "test_st/ring3",
			wantVal:  "test_string3",
		},
		{
			name:     "Positive test. Double slash",
			inputVal: "te//st_string4",
			wantVal:  "test_string4",
		},
	}

	for _, tt := range tests {
		// запускаем каждый тест
		resultStr := ""
		t.Run(tt.name, func(t *testing.T) {
			resultStr = trimSlashes(tt.inputVal)
			if resultStr != tt.wantVal {
				t.Errorf("TEST_ERROR: input value is %s, want is %s but result is %s", tt.inputVal, tt.wantVal, resultStr)
			}
		})
	}
}

func TestReductURL(t *testing.T) {
	var us = make(map[string]string)

	type inputStruct struct {
		url            string
		shortURLLength int
		urlStorage     map[string]string
	}

	tests := []struct {
		name           string
		inputs         inputStruct
		lenghtOfResult int
	}{ //Test table
		{
			name: "Positive test. Lenght of the shortenURL is 10",
			inputs: inputStruct{
				url:            "http://google.com/qwertyuiopasdfghjkkllzxcvbnm",
				shortURLLength: 10,
				urlStorage:     us,
			},
			lenghtOfResult: 10,
		},
		{
			name: "Positive test. Lenght of the shortenURL is 15",
			inputs: inputStruct{
				url:            "http://google.com/qwertyuiopasdfghjkkllzxcvbnm",
				shortURLLength: 15,
				urlStorage:     us,
			},
			lenghtOfResult: 15,
		},
		{
			name: "Positive test. URL with strange symbols",
			inputs: inputStruct{
				url:            "http://google.com/erty?ui&opa!@#$%^&*()_+~_sdfghjkkllzxcvbnm",
				shortURLLength: 15,
				urlStorage:     us,
			},
			lenghtOfResult: 15,
		},
		{
			name: "Positive test. Check adding token into urlStorage",
			inputs: inputStruct{
				url:            "http://google.com/erty?ui&opa!@#$%^&*()_+~_sdfghjkkllzxcvbnm",
				shortURLLength: 15,
				urlStorage:     us,
			},
			lenghtOfResult: 15,
		},
	}
	for _, tt := range tests {
		// запускаем каждый тест
		var resultStr string
		var takenURL string
		var ok bool

		t.Run(tt.name, func(t *testing.T) {
			resultStr = reductURL(tt.inputs.url, tt.inputs.shortURLLength, tt.inputs.urlStorage)
			fmt.Printf("TEST_DEBUG: Shortened token is '%s' for URL '%s'.\n", resultStr, tt.inputs.url)
			takenURL, ok = tt.inputs.urlStorage[resultStr]
			if len(resultStr) != tt.lenghtOfResult {
				t.Errorf("TEST_ERROR: input url is %s, wanted lenght of the resul string is %d but the outpus string is %s", tt.inputs.url, tt.lenghtOfResult, resultStr)
			} else if !ok {
				t.Errorf("TEST_ERROR: Token for URL '%s' didn't save into URL storage.\n", tt.inputs.url)
			} else if takenURL != tt.inputs.url {
				t.Errorf("TEST_ERROR: Gotten URL from the storage ('%s') doesn't match with input URL (%s).\n", resultStr, tt.inputs.url)
			}
		})
	}
}

func TestGetURLFromStorage(t *testing.T) {
	var us = make(map[string]string)

	us["qwerfadsfd"] = "https://golang-blog.blogspot.com/2020/01/map-golang.html"
	us["8rewq78rqew"] = "https://ru.wikipedia.org/wiki/%D0%9A%D0%B8%D1%80%D0%B8%D0%BB%D0%BB%D0%B8%D1%86%D0%B0"
	us["lahfsdafnb4121l"] = "https://ru.wikipedia.org/wiki/%D0%A3%D0%BC%D0%BB%D0%B0%D1%83%D1%82_(%D0%B4%D0%B8%D0%B0%D0%BA%D1%80%D0%B8%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B9_%D0%B7%D0%BD%D0%B0%D0%BA)"
	us["3123123123123"] = "https://en.wikipedia.org/wiki/Hungarian_alphabet"
	us["KJFASSFASDJSJ"] = "https://en.wikipedia.org/wiki/Latin_alphabet"

	type inputStruct struct {
		id         string
		urlStorage map[string]string
	}

	tests := []struct {
		name      string
		inputs    inputStruct
		resultStr string
	}{ //Test table
		{
			name: "Positive test 1. Get URL from URL storage (only letters).",
			inputs: inputStruct{
				id:         "qwerfadsfd",
				urlStorage: us,
			},
			resultStr: "https://golang-blog.blogspot.com/2020/01/map-golang.html",
		},
		{
			name: "Positive test 2. Get URL from URL storage (only digits).",
			inputs: inputStruct{
				id:         "3123123123123",
				urlStorage: us,
			},
			resultStr: "https://en.wikipedia.org/wiki/Hungarian_alphabet",
		},
		{
			name: "Positive test 3. Get URL from URL storage (long URL).",
			inputs: inputStruct{
				id:         "lahfsdafnb4121l",
				urlStorage: us,
			},
			resultStr: "https://ru.wikipedia.org/wiki/%D0%A3%D0%BC%D0%BB%D0%B0%D1%83%D1%82_(%D0%B4%D0%B8%D0%B0%D0%BA%D1%80%D0%B8%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B9_%D0%B7%D0%BD%D0%B0%D0%BA)",
		},
		{
			name: "Negative test 1. No such token.",
			inputs: inputStruct{
				id:         "afddsjdfsfasdf",
				urlStorage: us,
			},
			resultStr: "",
		},
	}
	for _, tt := range tests {
		// запускаем каждый тест
		var resultStr string
		t.Run(tt.name, func(t *testing.T) {
			resultStr = getURLFromStorage(tt.inputs.id, tt.inputs.urlStorage)
			fmt.Printf("TEST_DEBUG: For token '%s' returned URL is '%s'.\n", resultStr, tt.inputs.id)
			if resultStr != tt.resultStr {
				t.Errorf("TEST_ERROR: Returned  from storage string '%s'  for token '%s' doesn't match with stored one '%s'.\n", resultStr, tt.inputs.id, tt.resultStr)
			}
		})
	}
}

func TestPostAndGetHandler(t *testing.T) {

	URLStorage = make(map[string]string)
	URLStorage["qwerfadsfd"] = "https://golang-blog.blogspot.com/2020/01/map-golang.html"
	URLStorage["8rewq78rqew"] = "https://ru.wikipedia.org/wiki/%D0%9A%D0%B8%D1%80%D0%B8%D0%BB%D0%BB%D0%B8%D1%86%D0%B0"
	URLStorage["lahfsdafnb4121l"] = "https://ru.wikipedia.org/wiki/%D0%A3%D0%BC%D0%BB%D0%B0%D1%83%D1%82_(%D0%B4%D0%B8%D0%B0%D0%BA%D1%80%D0%B8%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B9_%D0%B7%D0%BD%D0%B0%D0%BA)"
	URLStorage["3123123123123"] = "https://en.wikipedia.org/wiki/Hungarian_alphabet"
	URLStorage["KJFASSFASDJSJ"] = "https://en.wikipedia.org/wiki/Latin_alphabet"

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
			} else if tt.input.reqMethod == http.MethodPost && len(string(resultBody)) != len(HostURL)+ShortURLLength {
				t.Errorf("TEST_ERROR: Wrong lenght of result (%d). Should be equal len(HostURL)+ShortURLLength (%d).\n", len(string(resultBody)), len(HostURL)+ShortURLLength)
			} else if tt.input.reqMethod == http.MethodGet && resultLocation != nil {
				if resultLocationFullURLStr != tt.want.respLocationHeaderStr {
					t.Errorf("TEST_ERROR: Location header is '%s' but should be '%s'.\n", string(resultLocationFullURLStr), tt.want.respLocationHeaderStr)
				}
			}
		})
	}
}
