package config

import (
	"fmt"
	"os"
	"testing"
)

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
			if tt.inputs.inputSP.HostSocketAddrStr != tt.wants.hostSocketAddrStr {
				t.Errorf("TEST_ERROR: Global var HostSocketAddrStr '%s' is not equal wants.hostSocketAddrSt '%s'.\n", tt.inputs.inputSP.HostSocketAddrStr, tt.wants.hostSocketAddrStr)
			} else if tt.inputs.inputSP.BaseURLStr != tt.wants.baseURLStr {
				t.Errorf("TEST_ERROR: Global var BaseURLStr '%s' is not equal wants.baseURLStr '%s'.\n", tt.inputs.inputSP.BaseURLStr, tt.wants.baseURLStr)
			} else if tt.inputs.inputSP.RepoFilePathStr != tt.wants.repoFilePathStr {
				t.Errorf("TEST_ERROR: Global var RepoFilePathStr '%s' is not equal wants.repoFilePathStr '%s'.\n", tt.inputs.inputSP.RepoFilePathStr, tt.wants.repoFilePathStr)
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
			if sp.HostSocketAddrStr != tt.wants.hostSocketStr {
				t.Errorf("TEST_ERROR: Global var HostSocketAddrStr '%s' is not equal wants.hostSocketStr '%s'.\n", sp.HostSocketAddrStr, tt.wants.hostSocketStr)
			} else if sp.BaseURLStr != tt.wants.baseURLStr {
				t.Errorf("TEST_ERROR: Global var BaseURLStr '%s' is not equal wants.BaseURLStr '%s'.\n", sp.BaseURLStr, tt.wants.baseURLStr)
			} else if sp.RepoFilePathStr != tt.wants.fileStoragePathStr {
				t.Errorf("TEST_ERROR: Global var RepoFilePathStr '%s' is not equal wants.fileStoragePathStr '%s'.\n", sp.RepoFilePathStr, tt.wants.fileStoragePathStr)
			}

		})
	}
}
