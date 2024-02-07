package memrepo

import (
	"fmt"
	"github.com/DmitrySkalnenkov/reduction/internal/entity"
	"testing"
)

func TestMemRepo_GetURLFromRepo(t *testing.T) {
	var ur MemRepo
	ur.InitRepo("")
	ur = MemRepo{
		urlMap: map[string]entity.URLUser{
			"qwerfadsfd":      {"https://golang-blog.blogspot.com/2020/01/map-golang.html", 0},
			"8rewq78rqew":     {"https://ru.wikipedia.org/wiki/%D0%9A%D0%B8%D1%80%D0%B8%D0%BB%D0%BB%D0%B8%D1%86%D0%B0", 0},
			"lahfsdafnb4121l": {"https://ru.wikipedia.org/wiki/%D0%A3%D0%BC%D0%BB%D0%B0%D1%83%D1%82_(%D0%B4%D0%B8%D0%B0%D0%BA%D1%80%D0%B8%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B9_%D0%B7%D0%BD%D0%B0%D0%BA)", 0},
			"3123123123123":   {"https://en.wikipedia.org/wiki/Hungarian_alphabet", 0},
			"KJFASSFASDJSJ":   {"https://en.wikipedia.org/wiki/Latin_alphabet", 0},
		},
	}

	type inputStruct struct {
		id         string
		urlStorage MemRepo
	}

	tests := []struct {
		name          string
		inputs        inputStruct
		wantResultStr string
		wantOk        bool
	}{ //Test table
		{
			name: "Positive test 1. Get URL from URL storage (only letters).",
			inputs: inputStruct{
				id:         "qwerfadsfd",
				urlStorage: ur,
			},
			wantResultStr: "https://golang-blog.blogspot.com/2020/01/map-golang.html",
			wantOk:        true,
		},
		{
			name: "Positive test 2. Get URL from URL storage (only digits).",
			inputs: inputStruct{
				id:         "3123123123123",
				urlStorage: ur,
			},
			wantResultStr: "https://en.wikipedia.org/wiki/Hungarian_alphabet",
			wantOk:        true,
		},
		{
			name: "Positive test 3. Get URL from URL storage (long URL).",
			inputs: inputStruct{
				id:         "lahfsdafnb4121l",
				urlStorage: ur,
			},
			wantResultStr: "https://ru.wikipedia.org/wiki/%D0%A3%D0%BC%D0%BB%D0%B0%D1%83%D1%82_(%D0%B4%D0%B8%D0%B0%D0%BA%D1%80%D0%B8%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B9_%D0%B7%D0%BD%D0%B0%D0%BA)",
			wantOk:        true,
		},
		{
			name: "Negative test 1. No such token.",
			inputs: inputStruct{
				id:         "afddsjdfsfasdf",
				urlStorage: ur,
			},
			wantResultStr: "",
			wantOk:        false,
		},
	}
	for _, tt := range tests {
		// запускаем каждый тест
		var resultStr string
		var ok bool
		t.Run(tt.name, func(t *testing.T) {
			resultStr, ok = tt.inputs.urlStorage.GetURLFromRepo(tt.inputs.id)
			fmt.Printf("TEST_DEBUG: For token '%s' returned URL is '%s'.\n", tt.inputs.id, resultStr)
			if resultStr != tt.wantResultStr || ok != tt.wantOk {
				t.Errorf("TEST_ERROR: Returned  from storage string '%s'  for token '%s' doesn't match with stored one '%s', or wantOk value (%t) unexpected.\n", resultStr, tt.inputs.id, tt.wantResultStr, ok)
			}
		})
	}
}

func TestMemRepo_SetURLIntoRepo(t *testing.T) {
	var ur MemRepo
	ur.InitRepo("")
	ur = MemRepo{
		urlMap: map[string]entity.URLUser{
			"qwerfadsfd":  {"https://golang-blog.blogspot.com/2020/01/map-golang.html", 0},
			"8rewq78rqew": {"https://ru.wikipedia.org/wiki/%D0%9A%D0%B8%D1%80%D0%B8%D0%BB%D0%BB%D0%B8%D1%86%D0%B0", 0},
		},
	}

	//ur.urlMap["qwerfadsfd"].URL = "https://golang-blog.blogspot.com/2020/01/map-golang.html"
	//ur.urlMap["8rewq78rqew"].URL = "https://ru.wikipedia.org/wiki/%D0%9A%D0%B8%D1%80%D0%B8%D0%BB%D0%BB%D0%B8%D1%86%D0%B0"

	type inputStruct struct {
		token      string
		value      string
		urlStorage MemRepo
	}

	tests := []struct {
		name   string
		inputs inputStruct
	}{ //Test table
		{
			name: "Positive test 1. Set URL into URL storage (only letters).",
			inputs: inputStruct{
				token:      "qwerfadsfd",
				value:      "https://go.dev/tour/moretypes/19",
				urlStorage: ur,
			},
		},
		{
			name: "Positive test 2. Set URL into URL storage (only digits).",
			inputs: inputStruct{
				token:      "41341414134",
				value:      "https://go.dev/doc/faq#pass_by_value",
				urlStorage: ur,
			},
		},
		{
			name: "Positive test 3. Empty string as token into URL storage (only digits).",
			inputs: inputStruct{
				token:      "",
				value:      "https://practicum.yandex.ru/learn/go-advanced-self-paced/courses/8bca0296-484d-45dc-b9ab-f01e0f44f9f4/sprints/145736/topics/63027ac1-f19b-405d-bad5-49e3bbddf30b/lessons/572d89a8-1713-457a-927a-90c2280757bc/",
				urlStorage: ur,
			},
		},
	}
	for _, tt := range tests {
		// запускаем каждый тест
		var isValueExist bool
		var resultURL string
		t.Run(tt.name, func(t *testing.T) {
			tt.inputs.urlStorage.SetURLIntoRepo(tt.inputs.token, tt.inputs.value)
			fmt.Printf("TEST_DEBUG: For SetURLIntoRepo set token '%s' with URL '%s' into URLStorage.\n", tt.inputs.token, tt.inputs.value)
			resultURL, isValueExist = tt.inputs.urlStorage.GetURLFromRepo(tt.inputs.token)
			if resultURL != tt.inputs.value || !isValueExist {
				t.Errorf("TEST_ERROR: URL '%s' for token '%s' doesn't set correctly into the URLStorage.\n", resultURL, tt.inputs.token)
			}
		})
	}
}