package app

import (
	"fmt"
	"github.com/DmitrySkalnenkov/reduction/internal/storage"
	"os"
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
	var ur storage.MemRepo
	ur.InitRepo()

	type inputStruct struct {
		url            string
		shortURLLength int
		urlStorage     *storage.MemRepo
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
			resultStr = ReduceURL(tt.inputs.url, tt.inputs.shortURLLength, tt.inputs.urlStorage)
			fmt.Printf("TEST_DEBUG: Shortened token is '%s' for URL '%s'.\n", resultStr, tt.inputs.url)
			takenURL, ok = tt.inputs.urlStorage.GetURLFromRepo(resultStr)
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

func TestGetEnvs(t *testing.T) {
	type inputStruct struct {
		hostSocketStr      string
		baseURLStr         string
		fileStoragePathStr string
	}

	type wantStruct struct {
		hostSocketStr      string
		baseURLStr         string
		fileStoragePathStr string
	}

	tests := []struct {
		name   string
		inputs inputStruct
		wants  wantStruct
	}{ //Test table
		{
			name: "Positive test 1. Usual address and URL (localhost)",
			inputs: inputStruct{
				hostSocketStr:      "127.0.0.1:8080",
				baseURLStr:         "http://google.com:5555",
				fileStoragePathStr: "/tmp/temp1.json",
			},
			wants: wantStruct{
				hostSocketStr:      "127.0.0.1:8080",
				baseURLStr:         "http://google.com:5555",
				fileStoragePathStr: "/tmp/temp1.json",
			},
		},
		{
			name: "Positive test 2. Usual address and URL (127.0.0.1)",
			inputs: inputStruct{
				hostSocketStr:      "localhost:9999",
				baseURLStr:         "http://yandex.ru",
				fileStoragePathStr: "/tmp/temp2.json",
			},
			wants: wantStruct{
				hostSocketStr:      "localhost:9999",
				baseURLStr:         "http://yandex.ru",
				fileStoragePathStr: "/tmp/temp2.json",
			},
		},
		{
			name: "Positive test 3. Empty address",
			inputs: inputStruct{
				hostSocketStr:      "",
				baseURLStr:         "http://yandex.ru",
				fileStoragePathStr: "/tmp/temp3.json",
			},
			wants: wantStruct{
				hostSocketStr:      "localhost:8080",
				baseURLStr:         "http://yandex.ru",
				fileStoragePathStr: "/tmp/temp3.json",
			},
		},
		{
			name: "Positive test 4. Empty BaseURL",
			inputs: inputStruct{
				hostSocketStr: "localhost:8090",
				baseURLStr:    "",
			},
			wants: wantStruct{
				hostSocketStr:      "localhost:8090",
				baseURLStr:         "http://localhost:8080",
				fileStoragePathStr: DefaultRepoFilePath,
			},
		},
	}
	for _, tt := range tests {
		// запускаем каждый тест
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("SERVER_ADDRESS", tt.inputs.hostSocketStr)
			os.Setenv("BASE_URL", tt.inputs.baseURLStr)
			os.Setenv("FILE_STORAGE_PATH", tt.inputs.fileStoragePathStr)
			var sp ServerParameters
			sp.GetEnvs()
			sp.CheckParamPriority()
			fmt.Printf("TEST_DEBUG: SERVER_ADDRESS enviroment variable is set to '%s'\n", tt.inputs.hostSocketStr)
			fmt.Printf("TEST_DEBUG: BASE_URL enviroment variable is set to '%s'\n", tt.inputs.baseURLStr)
			if HostSocketAddrStr != tt.wants.hostSocketStr {
				t.Errorf("TEST_ERROR: Global var HostSocketAddrStr '%s' is not equal wants.hostSocketStr '%s'.\n", HostSocketAddrStr, tt.wants.hostSocketStr)
			} else if BaseURLStr != tt.wants.baseURLStr {
				t.Errorf("TEST_ERROR: Global var BaseURLStr '%s' is not equal wants.BaseURLStr '%s'.\n", BaseURLStr, tt.wants.baseURLStr)
			} else if RepoFilePathStr != tt.wants.fileStoragePathStr {
				t.Errorf("TEST_ERROR: Global var RepoFilePathStr '%s' is not equal wants.fileStoragePathStr '%s'.\n", RepoFilePathStr, tt.wants.fileStoragePathStr)
			}

		})
	}
}

func TestCheckParamPriority(t *testing.T) {
	type inputStruct struct {
		inputSP ServerParameters
	}

	type wantStruct struct {
		hostSocketAddrStr string
		baseURLStr        string
		repoFilePathStr   string
	}

	tests := []struct {
		name   string
		inputs inputStruct
		wants  wantStruct
	}{ //Test table
		{
			name: "Positive test 1. (Usual address, URL and path for env)",
			inputs: inputStruct{
				inputSP: ServerParameters{
					envSocketAddr:    "127.0.1.1:8089",
					isEnvSocketAddr:  true,
					envBaseURL:       "http://local.com/path",
					isEnvBaseURL:     true,
					envRepoFilePath:  "/tmp/tmp.json",
					isRepoFilePath:   true,
					flagSocketAddr:   "",
					flagBaseURL:      "",
					flagRepoFilePath: "",
				},
			},
			wants: wantStruct{
				hostSocketAddrStr: "127.0.1.1:8089",
				baseURLStr:        "http://local.com/path",
				repoFilePathStr:   "/tmp/tmp.json",
			},
		},
		{
			name: "Positive test 2. (Env param is empty, flag exists)",
			inputs: inputStruct{
				inputSP: ServerParameters{
					envSocketAddr:    "",
					isEnvSocketAddr:  true,
					envBaseURL:       "",
					isEnvBaseURL:     true,
					envRepoFilePath:  "",
					isRepoFilePath:   true,
					flagSocketAddr:   "127.1.1.1:8888",
					flagBaseURL:      "https://gogol.ru",
					flagRepoFilePath: "/tmp/tmp1.json",
				},
			},
			wants: wantStruct{
				hostSocketAddrStr: "127.1.1.1:8888",
				baseURLStr:        "https://gogol.ru",
				repoFilePathStr:   "/tmp/tmp1.json",
			},
		},
		{
			name: "Positive test 3. (Envs don't set, flags exist)",
			inputs: inputStruct{
				inputSP: ServerParameters{
					envSocketAddr:    "",
					isEnvSocketAddr:  false,
					envBaseURL:       "",
					isEnvBaseURL:     false,
					envRepoFilePath:  "",
					isRepoFilePath:   false,
					flagSocketAddr:   "127.1.1.1:8888",
					flagBaseURL:      "https://gogol.ru",
					flagRepoFilePath: "/tmp/tmp1.json",
				},
			},
			wants: wantStruct{
				hostSocketAddrStr: "127.1.1.1:8888",
				baseURLStr:        "https://gogol.ru",
				repoFilePathStr:   "/tmp/tmp1.json",
			},
		},
		{
			name: "Positive test 4. (Envs don't set, flags don't set)",
			inputs: inputStruct{
				inputSP: ServerParameters{
					envSocketAddr:    "",
					isEnvSocketAddr:  false,
					envBaseURL:       "",
					isEnvBaseURL:     false,
					envRepoFilePath:  "",
					isRepoFilePath:   false,
					flagSocketAddr:   "",
					flagBaseURL:      "",
					flagRepoFilePath: "",
				},
			},
			wants: wantStruct{
				hostSocketAddrStr: DefaultSocketAddr,
				baseURLStr:        DefaultHostURL,
				repoFilePathStr:   DefaultRepoFilePath,
			},
		},
	}
	for _, tt := range tests {
		// запускаем каждый тест
		t.Run(tt.name, func(t *testing.T) {
			tt.inputs.inputSP.CheckParamPriority()
			if HostSocketAddrStr != tt.wants.hostSocketAddrStr {
				t.Errorf("TEST_ERROR: Global var HostSocketAddrStr '%s' is not equal wants.hostSocketAddrSt '%s'.\n", HostSocketAddrStr, tt.wants.hostSocketAddrStr)
			} else if BaseURLStr != tt.wants.baseURLStr {
				t.Errorf("TEST_ERROR: Global var BaseURLStr '%s' is not equal wants.baseURLStr '%s'.\n", BaseURLStr, tt.wants.baseURLStr)
			} else if RepoFilePathStr != tt.wants.repoFilePathStr {
				t.Errorf("TEST_ERROR: Global var RepoFilePathStr '%s' is not equal wants.repoFilePathStr '%s'.\n", RepoFilePathStr, tt.wants.repoFilePathStr)
			}
		})
	}
}
