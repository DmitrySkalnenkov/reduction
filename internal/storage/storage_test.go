package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestGetURLFromStorage(t *testing.T) {
	var ur MemRepo
	ur.InitRepo("")

	ur.urlMap["qwerfadsfd"] = "https://golang-blog.blogspot.com/2020/01/map-golang.html"
	ur.urlMap["8rewq78rqew"] = "https://ru.wikipedia.org/wiki/%D0%9A%D0%B8%D1%80%D0%B8%D0%BB%D0%BB%D0%B8%D1%86%D0%B0"
	ur.urlMap["lahfsdafnb4121l"] = "https://ru.wikipedia.org/wiki/%D0%A3%D0%BC%D0%BB%D0%B0%D1%83%D1%82_(%D0%B4%D0%B8%D0%B0%D0%BA%D1%80%D0%B8%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B9_%D0%B7%D0%BD%D0%B0%D0%BA)"
	ur.urlMap["3123123123123"] = "https://en.wikipedia.org/wiki/Hungarian_alphabet"
	ur.urlMap["KJFASSFASDJSJ"] = "https://en.wikipedia.org/wiki/Latin_alphabet"

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

func TestSetURLIntoStorage(t *testing.T) {
	var ur MemRepo
	ur.InitRepo("")

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
			if resultURL != tt.inputs.value || !isValueExist {
				t.Errorf("TEST_ERROR: URL '%s' for token '%s' doesn't set correctly into the URLStorage.\n", resultURL, tt.inputs.token)
			}
		})
	}
}

func TestFileRepo_InitRepo(t *testing.T) {
	var ur FileRepo
	filePath := "/tmp/temp_repository_file.json"
	ur.InitRepo(filePath)
	_, err := os.Stat(filePath)
	if err != nil {
		t.Errorf("TEST_ERROR: Cannot init file with path '%s'.\n", filePath)
	}
}

func TestFileRepo_PrintRepo(t *testing.T) {
	var ur FileRepo
	filePath := "/tmp/temp.json"
	ur.InitRepo(filePath)
	ur.PrintRepo()
}

func TestFileRepo_SetURLIntoRepo(t *testing.T) {
	var ur FileRepo
	filePath := "/tmp/temp_test_for_set_url.json"
	ur.InitRepo(filePath)
	jr1 := jsonRepo{
		jsonSlice: []jsonLine{
			{
				Token: "qwer",
				URL:   "http://google.com/1234",
			},
			{
				Token: "asdf",
				URL:   "http://yandex.com/1234",
			},
			{
				Token: "zxcv",
				URL:   "http://altavista.com/1234",
			},
			{
				Token: "3123123123123",
				URL:   "https://en.wikipedia.org/wiki/Hungarian_alphabet",
			},
			{
				Token: "lahfsdafnb4121l",
				URL:   "https://ru.wikipedia.org/wiki/%D0%A3%D0%BC%D0%BB%D0%B0%D1%83%D1%82_(%D0%B4%D0%B8%D0%B0%D0%BA%D1%80%D0%B8%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B9_%D0%B7%D0%BD%D0%B0%D0%BA)",
			},
		},
	}
	DumpRepoToJSONFile(&jr1, filePath)

	type inputStruct struct {
		token      string
		value      string
		urlStorage FileRepo
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

func TestFileRepo_GetURLFromRepo(t *testing.T) {
	var ur FileRepo
	filePath := "/tmp/temp_test_for_get_url.json"
	ur.InitRepo(filePath)
	jr1 := jsonRepo{
		jsonSlice: []jsonLine{
			{
				Token: "qwer",
				URL:   "http://google.com/1234",
			},
			{
				Token: "asdf",
				URL:   "http://yandex.com/1234",
			},
			{
				Token: "zxcv",
				URL:   "http://altavista.com/1234",
			},
			{
				Token: "3123123123123",
				URL:   "https://en.wikipedia.org/wiki/Hungarian_alphabet",
			},
			{
				Token: "lahfsdafnb4121l",
				URL:   "https://ru.wikipedia.org/wiki/%D0%A3%D0%BC%D0%BB%D0%B0%D1%83%D1%82_(%D0%B4%D0%B8%D0%B0%D0%BA%D1%80%D0%B8%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B9_%D0%B7%D0%BD%D0%B0%D0%BA)",
			},
		},
	}
	DumpRepoToJSONFile(&jr1, filePath)
	type inputStruct struct {
		id         string
		urlStorage FileRepo
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
				id:         "asdf",
				urlStorage: ur,
			},
			wantResultStr: "http://yandex.com/1234",
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

func TestDumpRepoToJSONFile(t *testing.T) {
	jr1 := jsonRepo{
		jsonSlice: []jsonLine{
			{
				Token: "qwer",
				URL:   "http://google.com/1234",
			},
			{
				Token: "asdf",
				URL:   "http://yandex.com/1234",
			},
			{
				Token: "zxcv",
				URL:   "http://altavista.com/1234",
			},
		},
	}
	var jr2 jsonRepo
	filePath := "/tmp/temp_test_for_dump.json"
	DumpRepoToJSONFile(&jr1, filePath)
	jsonRepoFile, err := os.OpenFile(filePath, os.O_RDONLY, 0777)
	if err != nil {
		fmt.Printf("ERROR: Filed to open the repository file '%s'.\n", filePath)
	}
	defer jsonRepoFile.Close()
	fromFile, err := io.ReadAll(jsonRepoFile)
	if err != nil {
		fmt.Printf("ERROR: Cannot read data from repo file '%s'.\n", filePath)
	}
	err = json.Unmarshal(fromFile, &jr2.jsonSlice)
	if err == nil {
		fmt.Printf("INFO: Data from repository file were restored succesfully.\n")
	} else {
		fmt.Printf("ERROR: '%s'. Data from file '%s' cannot be restored.\n", err, filePath)
	}

	for i := 0; i < len(jr1.jsonSlice); i++ {
		if jr1.jsonSlice[i] != jr2.jsonSlice[i] {
			t.Errorf("TEST_ERROR: Dumped and restored stucts jsonRepo don't equal.\n Dumped: %v \n Restored: %v\n", jr1, jr2)
			return
		}
	}
	os.Remove(filePath)
}

func TestRestoreRepoFromJSONFile(t *testing.T) {
	var toFile []byte
	jr1 := jsonRepo{
		jsonSlice: []jsonLine{
			{
				Token: "qwer",
				URL:   "http://google.com/1234",
			},
			{
				Token: "asdf",
				URL:   "http://yandex.com/1234",
			},
			{
				Token: "zxcv",
				URL:   "http://altavista.com/1234",
			},
		},
	}
	var jr2 jsonRepo
	filePath := "/tmp/temp_test_for_restore.json"
	jsonRepoFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Printf("ERROR: Filed to open the repository file '%s'.\n", filePath)
	}
	defer jsonRepoFile.Close()
	if len(jr1.jsonSlice) > 0 {
		toFile, err = json.Marshal(jr1.jsonSlice)
		if err != nil {
			fmt.Printf("ERROR: Cannot marshal repo.urlMap '%v' to JSON.\n", jr1)
		}
		jsonRepoFile.Truncate(0)
		jsonRepoFile.Seek(0, 0)
		_, err = jsonRepoFile.WriteString(string(toFile))
		if err == nil {
			fmt.Printf("INFO: URL repository dumped into the file '%s'.\n", filePath)
		} else {
			fmt.Printf("ERROR: Cannot dump JSON string '%s' to file '%s'.\n", string(toFile), filePath)
		}
		RestoreRepoFromJSONFile(&jr2, filePath)

	}
	for i := 0; i < len(jr1.jsonSlice); i++ {
		if jr1.jsonSlice[i] != jr2.jsonSlice[i] {
			t.Errorf("TEST_ERROR: Dumped and restored stucts jsonRepo don't equal.\n Dumped: %v \n Restored: %v\n", jr1, jr2)
			return
		}
	}
	os.Remove(filePath)
}

func TestURLStorageInit(t *testing.T) {

	tests := []struct {
		name        string
		filePath    string
		wantTypeStr string
	}{ //Test table
		{
			name:        "Positive test 1. Memory repository.",
			filePath:    "",
			wantTypeStr: "*storage.MemRepo",
		},
		{
			name:        "Positive test 1. Memory repository.",
			filePath:    "/tmp/temp_for_URLStorageInit.txt",
			wantTypeStr: "*storage.FileRepo",
		},
	}
	for _, tt := range tests {
		// запускаем каждый тест
		t.Run(tt.name, func(t *testing.T) {
			us := URLStorageInit(tt.filePath)
			curTypeStr := fmt.Sprintf("%T", us)
			fmt.Printf("Current type is '%s'\n", curTypeStr)
			if curTypeStr != tt.wantTypeStr {
				t.Errorf("TEST_ERROR: Type of storage ('%s') doesn't match expected ('%s').\n", curTypeStr, tt.wantTypeStr)
			}
			if tt.filePath != "" {
				os.Remove(tt.filePath)
			}
		})
	}
}
