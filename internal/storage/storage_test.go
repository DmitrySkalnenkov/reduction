package storage

import (
	"fmt"

	"testing"
)

func TestGetURLFromStorage(t *testing.T) {
	//var us = make(map[string]string)
	var ur Repository
	ur.urlMap = make(map[string]string)

	ur.urlMap["qwerfadsfd"] = "https://golang-blog.blogspot.com/2020/01/map-golang.html"
	ur.urlMap["8rewq78rqew"] = "https://ru.wikipedia.org/wiki/%D0%9A%D0%B8%D1%80%D0%B8%D0%BB%D0%BB%D0%B8%D1%86%D0%B0"
	ur.urlMap["lahfsdafnb4121l"] = "https://ru.wikipedia.org/wiki/%D0%A3%D0%BC%D0%BB%D0%B0%D1%83%D1%82_(%D0%B4%D0%B8%D0%B0%D0%BA%D1%80%D0%B8%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B9_%D0%B7%D0%BD%D0%B0%D0%BA)"
	ur.urlMap["3123123123123"] = "https://en.wikipedia.org/wiki/Hungarian_alphabet"
	ur.urlMap["KJFASSFASDJSJ"] = "https://en.wikipedia.org/wiki/Latin_alphabet"

	type inputStruct struct {
		id string
		//urlStorage map[string]string
		urlStorage Repository
	}

	tests := []struct {
		name      string
		inputs    inputStruct
		resultStr string
		ok        bool
	}{ //Test table
		{
			name: "Positive test 1. Get URL from URL storage (only letters).",
			inputs: inputStruct{
				id:         "qwerfadsfd",
				urlStorage: ur,
			},
			resultStr: "https://golang-blog.blogspot.com/2020/01/map-golang.html",
			ok:        true,
		},
		{
			name: "Positive test 2. Get URL from URL storage (only digits).",
			inputs: inputStruct{
				id:         "3123123123123",
				urlStorage: ur,
			},
			resultStr: "https://en.wikipedia.org/wiki/Hungarian_alphabet",
			ok:        true,
		},
		{
			name: "Positive test 3. Get URL from URL storage (long URL).",
			inputs: inputStruct{
				id:         "lahfsdafnb4121l",
				urlStorage: ur,
			},
			resultStr: "https://ru.wikipedia.org/wiki/%D0%A3%D0%BC%D0%BB%D0%B0%D1%83%D1%82_(%D0%B4%D0%B8%D0%B0%D0%BA%D1%80%D0%B8%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B9_%D0%B7%D0%BD%D0%B0%D0%BA)",
			ok:        true,
		},
		{
			name: "Negative test 1. No such token.",
			inputs: inputStruct{
				id:         "afddsjdfsfasdf",
				urlStorage: ur,
			},
			resultStr: "",
			ok:        false,
		},
	}
	for _, tt := range tests {
		// запускаем каждый тест
		var resultStr string
		var ok bool
		t.Run(tt.name, func(t *testing.T) {
			resultStr, ok = tt.inputs.urlStorage.GetURLFromStorage(tt.inputs.id)
			fmt.Printf("TEST_DEBUG: For token '%s' returned URL is '%s'.\n", resultStr, tt.inputs.id)
			if resultStr != tt.resultStr && ok != tt.ok {
				t.Errorf("TEST_ERROR: Returned  from storage string '%s'  for token '%s' doesn't match with stored one '%s', or ok value (%t) unexpected.\n", resultStr, tt.inputs.id, tt.resultStr, ok)
			}
		})
	}
}
