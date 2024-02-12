package services

import (
	"fmt"
	"github.com/DmitrySkalnenkov/reduction/internal/interfaces"
	"github.com/DmitrySkalnenkov/reduction/internal/models"
	"github.com/DmitrySkalnenkov/reduction/internal/repo/memrepo"
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

func TestReductURL(t *testing.T) {
	us := new(memrepo.MemRepo)
	us.InitRepo("")

	type inputStruct struct {
		urluser        models.URLUser
		shortURLLength int
		urlStorage     interfaces.DataRepo
	}

	tests := []struct {
		name           string
		inputs         inputStruct
		lenghtOfResult int
	}{ //Test table
		{
			name: "Positive test. Lenght of the shortenURL is 10",
			inputs: inputStruct{
				urluser:        models.URLUser{URL: "http://google.com/qwertyuiopasdfghjkkllzxcvbnm", UserID: 0},
				shortURLLength: 10,
				urlStorage:     us,
			},
			lenghtOfResult: 10,
		},
		{
			name: "Positive test. Lenght of the shortenURL is 15",
			inputs: inputStruct{
				urluser:        models.URLUser{URL: "http://google.com/qwertyuiopasdfghjkkllzxcvbnm", UserID: 0},
				shortURLLength: 15,
				urlStorage:     us,
			},
			lenghtOfResult: 15,
		},
		{
			name: "Positive test. URL with strange symbols",
			inputs: inputStruct{
				urluser:        models.URLUser{URL: "http://google.com/erty?ui&opa!@#$%^&*()_+~_sdfghjkkllzxcvbnm", UserID: 0},
				shortURLLength: 15,
				urlStorage:     us,
			},
			lenghtOfResult: 15,
		},
		{
			name: "Positive test. Check adding token into urlStorage",
			inputs: inputStruct{
				urluser:        models.URLUser{URL: "http://google.com/erty?ui&opa!@#$%^&*()_+~_sdfghjkkllzxcvbnm", UserID: 0},
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
			resultStr = ReduceURL(tt.inputs.urluser, tt.inputs.shortURLLength, tt.inputs.urlStorage)
			fmt.Printf("TEST_DEBUG: Shortened token is '%s' for URL '%s'.\n", resultStr, tt.inputs.urluser.URL)
			takenURL, ok = tt.inputs.urlStorage.GetURLFromRepo(resultStr)
			if len(resultStr) != tt.lenghtOfResult {
				t.Errorf("TEST_ERROR: input url is %s, wanted lenght of the resul string is %d but the outpus string is %s", tt.inputs.urluser.URL, tt.lenghtOfResult, resultStr)
			} else if !ok {
				t.Errorf("TEST_ERROR: Token for URL '%s' didn't save into URL storage.\n", tt.inputs.urluser.URL)
			} else if takenURL != tt.inputs.urluser.URL {
				t.Errorf("TEST_ERROR: Gotten URL from the storage ('%s') doesn't match with input URL (%s).\n", resultStr, tt.inputs.urluser.URL)
			}
		})
	}
}
