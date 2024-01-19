package app

import (
	"fmt"
	"github.com/DmitrySkalnenkov/reduction/internal/storage"
	"testing"
)

// (i2) Покройте сервис юнит-тестами. Сконцентрируйтесь на покрытии тестами эндпоинтов, чтобы защитить API сервиса от случайных изменений.
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
			resultStr = TrimSlashes(tt.inputVal)
			if resultStr != tt.wantVal {
				t.Errorf("TEST_ERROR: input value is %s, want is %s but result is %s", tt.inputVal, tt.wantVal, resultStr)
			}
		})
	}
}

func TestReductURL(t *testing.T) {
	//var us = make(map[string]string)
	var ur storage.Repository
	ur.Init()

	type inputStruct struct {
		url            string
		shortURLLength int
		urlStorage     *storage.Repository
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
				urlStorage:     &ur,
			},
			lenghtOfResult: 10,
		},
		{
			name: "Positive test. Lenght of the shortenURL is 15",
			inputs: inputStruct{
				url:            "http://google.com/qwertyuiopasdfghjkkllzxcvbnm",
				shortURLLength: 15,
				urlStorage:     &ur,
			},
			lenghtOfResult: 15,
		},
		{
			name: "Positive test. URL with strange symbols",
			inputs: inputStruct{
				url:            "http://google.com/erty?ui&opa!@#$%^&*()_+~_sdfghjkkllzxcvbnm",
				shortURLLength: 15,
				urlStorage:     &ur,
			},
			lenghtOfResult: 15,
		},
		{
			name: "Positive test. Check adding token into urlStorage",
			inputs: inputStruct{
				url:            "http://google.com/erty?ui&opa!@#$%^&*()_+~_sdfghjkkllzxcvbnm",
				shortURLLength: 15,
				urlStorage:     &ur,
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
			resultStr = ReductURL(tt.inputs.url, tt.inputs.shortURLLength, *tt.inputs.urlStorage)
			fmt.Printf("TEST_DEBUG: Shortened token is '%s' for URL '%s'.\n", resultStr, tt.inputs.url)
			takenURL, ok = tt.inputs.urlStorage.GetURLFromStorage(resultStr)
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
