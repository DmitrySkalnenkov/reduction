package filerepo

import (
	"encoding/json"
	"fmt"
	"github.com/DmitrySkalnenkov/reduction/internal/models"
	"io"
	"os"
	"testing"
)

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
	jr1 := models.JSONRepo{
		JSONSlice: []models.JSONLine{
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
		urluser    models.URLUser
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
				urluser:    models.URLUser{URL: "https://go.dev/tour/moretypes/19", UserID: 0},
				urlStorage: ur,
			},
		},
		{
			name: "Positive test 2. Set URL into URL storage (only digits).",
			inputs: inputStruct{
				token:      "41341414134",
				urluser:    models.URLUser{URL: "https://go.dev/doc/faq#pass_by_value", UserID: 0},
				urlStorage: ur,
			},
		},
		{
			name: "Positive test 3. Empty string as token into URL storage (only digits).",
			inputs: inputStruct{
				token:      "",
				urluser:    models.URLUser{URL: "https://practicum.yandex.ru/learn/go-advanced-self-paced/courses/8bca0296-484d-45dc-b9ab-f01e0f44f9f4/sprints/145736/topics/63027ac1-f19b-405d-bad5-49e3bbddf30b/lessons/572d89a8-1713-457a-927a-90c2280757bc/", UserID: 0},
				urlStorage: ur,
			},
		},
	}
	for _, tt := range tests {
		// запускаем каждый тест
		var isValueExist bool
		var resultURL string
		t.Run(tt.name, func(t *testing.T) {
			tt.inputs.urlStorage.SetURLIntoRepo(tt.inputs.token, tt.inputs.urluser.URL)
			fmt.Printf("TEST_DEBUG: For SetURLIntoRepo set token '%s' with URL '%s' into URLStorage.\n", tt.inputs.token, tt.inputs.urluser.URL)
			resultURL, isValueExist = tt.inputs.urlStorage.GetURLFromRepo(tt.inputs.token)
			if resultURL != tt.inputs.urluser.URL || !isValueExist {
				t.Errorf("TEST_ERROR: URL '%s' for token '%s' doesn't set correctly into the URLStorage.\n", resultURL, tt.inputs.token)
			}
		})
	}
}

func TestFileRepo_GetURLFromRepo(t *testing.T) {
	var ur FileRepo
	filePath := "/tmp/temp_test_for_get_url.json"
	ur.InitRepo(filePath)
	jr1 := models.JSONRepo{
		JSONSlice: []models.JSONLine{
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
				token:      "asdf",
				urlStorage: ur,
			},
			wantResultStr: "http://yandex.com/1234",
			wantOk:        true,
		},
		{
			name: "Positive test 2. Get URL from URL storage (only digits).",
			inputs: inputStruct{
				token:      "3123123123123",
				urlStorage: ur,
			},
			wantResultStr: "https://en.wikipedia.org/wiki/Hungarian_alphabet",
			wantOk:        true,
		},
		{
			name: "Positive test 3. Get URL from URL storage (long URL).",
			inputs: inputStruct{
				token:      "lahfsdafnb4121l",
				urlStorage: ur,
			},
			wantResultStr: "https://ru.wikipedia.org/wiki/%D0%A3%D0%BC%D0%BB%D0%B0%D1%83%D1%82_(%D0%B4%D0%B8%D0%B0%D0%BA%D1%80%D0%B8%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B9_%D0%B7%D0%BD%D0%B0%D0%BA)",
			wantOk:        true,
		},
		{
			name: "Negative test 1. No such token.",
			inputs: inputStruct{
				token:      "afddsjdfsfasdf",
				urlStorage: ur,
			},
			wantResultStr: "",
			wantOk:        false,
		},
	}
	for _, tt := range tests {
		// запускаем каждый тест
		var resultURL string
		var ok bool
		t.Run(tt.name, func(t *testing.T) {
			resultURL, ok = tt.inputs.urlStorage.GetURLFromRepo(tt.inputs.token)
			fmt.Printf("TEST_DEBUG: For token '%s' returned URL is '%s'.\n", tt.inputs.token, resultURL)
			if resultURL != tt.wantResultStr || ok != tt.wantOk {
				t.Errorf("TEST_ERROR: Returned  from storage string '%s'  for token '%s' doesn't match with stored one '%s', or wantOk urluser (%t) unexpected.\n", resultURL, tt.inputs.token, tt.wantResultStr, ok)
			}
		})
	}
}

func TestDumpRepoToJSONFile(t *testing.T) {
	jr1 := models.JSONRepo{
		JSONSlice: []models.JSONLine{
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
	var jr2 models.JSONRepo
	filePath := "/tmp/temp_test_for_dump.json"
	DumpRepoToJSONFile(&jr1, filePath)
	jsonRepoFile, err := os.OpenFile(filePath, os.O_RDONLY, 0777)
	if err != nil {
		fmt.Printf("ERROR: Filed to open the repository file '%s'.\n", filePath)
	}
	defer jsonRepoFile.Close()
	fromFile, err := io.ReadAll(jsonRepoFile)
	if err != nil {
		fmt.Printf("ERROR: Cannot read data from memrepo file '%s'.\n", filePath)
	}
	err = json.Unmarshal(fromFile, &jr2.JSONSlice)
	if err == nil {
		fmt.Printf("INFO: Data from repository file were restored succesfully.\n")
	} else {
		fmt.Printf("ERROR: '%s'. Data from file '%s' cannot be restored.\n", err, filePath)
	}

	for i := 0; i < len(jr1.JSONSlice); i++ {
		if jr1.JSONSlice[i] != jr2.JSONSlice[i] {
			t.Errorf("TEST_ERROR: Dumped and restored stucts jsonRepo don't equal.\n Dumped: %v \n Restored: %v\n", jr1, jr2)
			return
		}
	}
	os.Remove(filePath)
}

func TestRestoreRepoFromJSONFile(t *testing.T) {
	var toFile []byte
	jr1 := models.JSONRepo{
		JSONSlice: []models.JSONLine{
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
	var jr2 models.JSONRepo
	filePath := "/tmp/temp_test_for_restore.json"
	jsonRepoFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Printf("ERROR: Filed to open the repository file '%s'.\n", filePath)
	}
	defer jsonRepoFile.Close()
	if len(jr1.JSONSlice) > 0 {
		toFile, err = json.Marshal(jr1.JSONSlice)
		if err != nil {
			fmt.Printf("ERROR: Cannot marshal memrepo.urlMap '%v' to JSON.\n", jr1)
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
	for i := 0; i < len(jr1.JSONSlice); i++ {
		if jr1.JSONSlice[i] != jr2.JSONSlice[i] {
			t.Errorf("TEST_ERROR: Dumped and restored stucts jsonRepo don't equal.\n Dumped: %v \n Restored: %v\n", jr1, jr2)
			return
		}
	}
	os.Remove(filePath)
}
