package storage

import (
	"fmt"
	"os"

	"testing"
)

func TestGetURLFromStorage(t *testing.T) {
	var ur MemRepo
	ur.InitRepo()

	ur.urlMap["qwerfadsfd"] = "https://golang-blog.blogspot.com/2020/01/map-golang.html"
	ur.urlMap["8rewq78rqew"] = "https://ru.wikipedia.org/wiki/%D0%9A%D0%B8%D1%80%D0%B8%D0%BB%D0%BB%D0%B8%D1%86%D0%B0"
	ur.urlMap["lahfsdafnb4121l"] = "https://ru.wikipedia.org/wiki/%D0%A3%D0%BC%D0%BB%D0%B0%D1%83%D1%82_(%D0%B4%D0%B8%D0%B0%D0%BA%D1%80%D0%B8%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B9_%D0%B7%D0%BD%D0%B0%D0%BA)"
	ur.urlMap["3123123123123"] = "https://en.wikipedia.org/wiki/Hungarian_alphabet"
	ur.urlMap["KJFASSFASDJSJ"] = "https://en.wikipedia.org/wiki/Latin_alphabet"

	type inputStruct struct {
		id string
		//urlStorage map[string]string
		urlStorage MemRepo
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
			resultStr, ok = tt.inputs.urlStorage.GetURLFromRepo(tt.inputs.id)
			fmt.Printf("TEST_DEBUG: For token '%s' returned URL is '%s'.\n", resultStr, tt.inputs.id)
			if resultStr != tt.resultStr && ok != tt.ok {
				t.Errorf("TEST_ERROR: Returned  from storage string '%s'  for token '%s' doesn't match with stored one '%s', or ok value (%t) unexpected.\n", resultStr, tt.inputs.id, tt.resultStr, ok)
			}
		})
	}
}

func TestSetURLIntoStorage(t *testing.T) {
	var ur MemRepo
	ur.InitRepo()

	ur.urlMap["qwerfadsfd"] = "https://golang-blog.blogspot.com/2020/01/map-golang.html"
	ur.urlMap["8rewq78rqew"] = "https://ru.wikipedia.org/wiki/%D0%9A%D0%B8%D1%80%D0%B8%D0%BB%D0%BB%D0%B8%D1%86%D0%B0"

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
			if resultURL != tt.inputs.value && !isValueExist {
				t.Errorf("TEST_ERROR: URL '%s' for token '%s' doesn't set correctly into the URLStorage.\n", resultURL, tt.inputs.token)
			}
		})
	}
}

func TestDumpRepositoryToJSONFile(t *testing.T) {
	var ur1 MemRepo
	var ur2 MemRepo
	ur1.InitRepo()
	ur2.InitRepo()
	ur1.urlMap["qwerfadsfd"] = "https://golang-blog.blogspot.com/2020/01/map-golang.html"
	ur1.urlMap["8rewq78rqew"] = "https://ru.wikipedia.org/wiki/%D0%9A%D0%B8%D1%80%D0%B8%D0%BB%D0%BB%D0%B8%D1%86%D0%B0"
	ur1.urlMap["lahfsdafnb4121l"] = "https://ru.wikipedia.org/wiki/%D0%A3%D0%BC%D0%BB%D0%B0%D1%83%D1%82_(%D0%B4%D0%B8%D0%B0%D0%BA%D1%80%D0%B8%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B9_%D0%B7%D0%BD%D0%B0%D0%BA)"
	ur1.urlMap["3123123123123"] = "https://en.wikipedia.org/wiki/Hungarian_alphabet"
	ur1.urlMap["KJFASSFASDJSJ"] = "https://en.wikipedia.org/wiki/Latin_alphabet"

	type inputStruct struct {
		filePath   string
		urlStorage MemRepo
	}
	type wantStruct struct {
		urlStorage MemRepo
	}

	tests := []struct {
		name   string
		inputs inputStruct
	}{ //Test table
		{
			name: "Positive test.",
			inputs: inputStruct{
				filePath:   "/tmp/temp_repository_file.txt",
				urlStorage: ur1,
			},
		},
	}
	for _, tt := range tests {
		// запускаем каждый тест
		t.Run(tt.name, func(t *testing.T) {
			tt.inputs.urlStorage.PrintRepo()
			tt.inputs.urlStorage.DumpRepositoryToJSONFile(tt.inputs.filePath)
			fmt.Printf("TEST_DEBUG: MemRepo was dumped into the file '%s'.\n", tt.inputs.filePath)
			ur2.PrintRepo()
			ur2.RestoreRepositoryFromJSONFile(tt.inputs.filePath)
			ur2.PrintRepo()
		})
	}
}

func TestFileRepo_InitRepo(t *testing.T) {
	var ur FileRepo
	filePath := "/tmp/testFileRepo.txt"
	ur.InitRepo(filePath)
	_, err := os.Stat(filePath)
	if err != nil {
		t.Errorf("TEST_ERROR: Cannot init file with path '%s'.\n", filePath)
	}
}

func TestFileRepo_PrintRepo(t *testing.T) {
	var ur FileRepo
	filePath := "/tmp/testFileRepo.txt"
	ur.InitRepo(filePath)
	ur.PrintRepo()
}
