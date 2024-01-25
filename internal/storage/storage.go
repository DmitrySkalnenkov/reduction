package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

//(i2) Cлой storage должен имплементировать интерфейс хранения, например repositories. Это понадобится вам для подмены хранилища моком в тестах и использования DI.

type TxJSONMessage struct {
	URL string `json:"url"`
}
type RxJSONMessage struct {
	Result string `json:"result"`
}

type jsonLine struct {
	Token string `json:"token"`
	URL   string `json:"url"`
}

type jsonRepo struct {
	JSONSlice []jsonLine
}

type MemRepo struct {
	urlMap map[string]string
}

type FileRepo struct {
	urlFile     *os.File
	urlFilePath string
}

type Keeper interface {
	GetURLFromRepo(token string) (string, bool)
	SetURLIntoRepo(token string, value string)
	InitRepo(repoPath string)
	PrintRepo()
}

var URLStorage Keeper //MemRepo

// Return saved long URL from URL storage
func (repo *MemRepo) GetURLFromRepo(token string) (string, bool) {
	url, ok := repo.urlMap[token]
	if ok {
		fmt.Printf("DEBUG: Found long URL with token '%s' is URL storage.\n", token)
		return url, ok
	} else {
		fmt.Printf("DEBUG: Shortened URL with token '%s' not found in URL storage.\n", token)
		return "", ok
	}
}

// Set value (long URL) into repository for this token(shortened URL)
func (repo *MemRepo) SetURLIntoRepo(token string, value string) {
	repo.urlMap[token] = value
}

// Init()Initialization of MemRepo object
func (repo *MemRepo) InitRepo(repoPath string) {
	repo.urlMap = make(map[string]string)
}

func (repo *MemRepo) PrintRepo() {
	fmt.Println("VVVVVVVVVVVVVVVVVVVVVVVVVV")
	fmt.Println("DEBUG: UrlStorage. Begin:")
	for k, v := range repo.urlMap {
		fmt.Println(k, "value is", v)
	}
	fmt.Println("DEBUG: UrlStorage. End.")
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^")
}

// InitRepo() Initialization file repository. Create or trucate file with filePath if it exists.
func (repo *FileRepo) InitRepo(filePath string) {
	if filePath != "" {
		var err error
		repo.urlFilePath = filePath
		repo.urlFile, err = os.OpenFile(repo.urlFilePath, os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
			fmt.Printf("ERROR: Cannot open repository file '%s' for dumping .\n", filePath)
		}
		defer repo.urlFile.Close()
		return
	} else {
		fmt.Printf("ERROR: File path doesn't set.\n")
		return
	}
}

// PrintRepo() print content of file with filePath
func (repo *FileRepo) PrintRepo() {
	var err error
	if repo.urlFilePath != "" {
		repo.urlFile, err = os.OpenFile(repo.urlFilePath, os.O_RDONLY, 0777)
		if err != nil {
			fmt.Printf("ERROR: Cannot open repository file '%s' for reading.\n", repo.urlFilePath)
		}
		defer repo.urlFile.Close()
		dataFromFile, err := io.ReadAll(repo.urlFile)
		if err != nil {
			fmt.Printf("ERROR: Filed get data from repo file.\n")
			return
		}
		var jsonData jsonRepo
		err = json.Unmarshal(dataFromFile, &jsonData)
		if err != nil {
			fmt.Printf("ERROR: Filed unmarshal data from repo file.\n")
			return
		} else {
			for i := 0; i < len(jsonData.JSONSlice); i++ {
				fmt.Printf("'%s':'%s'", jsonData.JSONSlice[i].Token, jsonData.JSONSlice[i].URL)
			}
		}
	} else {
		fmt.Printf("ERROR: FileRepo path is empty. Cannot be used filerepository.\n")
		return
	}
}

// SetURLIntoRepo() Save token and url in JSON foramat into JSON file
func (repo *FileRepo) SetURLIntoRepo(token string, value string) {
	if repo.urlFilePath != "" {
		var tmpJSONRepo jsonRepo
		RestoreRepoFromJSONFile(&tmpJSONRepo, repo.urlFilePath)
		curJSONLine := jsonLine{Token: token, URL: value}
		for i := 0; i < len(tmpJSONRepo.JSONSlice); i++ {
			if curJSONLine.Token == tmpJSONRepo.JSONSlice[i].Token {
				fmt.Printf("INFO: Token with such value ('%s') already in JSON file repository. Token should be unique. The repo didn'token change.\n", curJSONLine.Token)
				return
			}
		}
		tmpJSONRepo.JSONSlice = append(tmpJSONRepo.JSONSlice, curJSONLine)
		DumpRepoToJSONFile(&tmpJSONRepo, repo.urlFilePath)
	}
}

// GetURLFromRepo() Get URL from file in JSON format. If token exists isExists = true
func (repo *FileRepo) GetURLFromRepo(token string) (string, bool) {
	var isURLExists = false
	var curURL = ""
	if repo.urlFilePath != "" {
		var tmpJSONRepo jsonRepo
		RestoreRepoFromJSONFile(&tmpJSONRepo, repo.urlFilePath)
		for i := 0; i < len(tmpJSONRepo.JSONSlice); i++ {
			if token == tmpJSONRepo.JSONSlice[i].Token {
				fmt.Printf("INFO: Token ('%s') was found in JSON file repository\n", token)
				isURLExists = true
				curURL = tmpJSONRepo.JSONSlice[i].URL
				return curURL, isURLExists
			}
		}
	}
	return curURL, isURLExists
}

// DumpRepoToJSONFIle() dumps data form jr to JSON file with filePath
func DumpRepoToJSONFile(jr *jsonRepo, filePath string) {
	if filePath != "" {
		var toFile []byte
		repoFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
		if err != nil {
			fmt.Printf("ERROR: Cannot open JSON repository file '%s' for dumping JSON-data.\n", filePath)
		}
		defer repoFile.Close()
		if len(jr.JSONSlice) > 0 {
			toFile, err = json.Marshal(jr.JSONSlice)
			if err != nil {
				fmt.Printf("ERROR: Cannot marshal repo.urlMap '%v' to JSON.\n", jr)
			}
			repoFile.Truncate(0)
			repoFile.Seek(0, 0)
			_, err = repoFile.WriteString(string(toFile))
			if err == nil {
				fmt.Printf("INFO: URL repository dumped into the file '%s'.\n", filePath)
			} else {
				fmt.Printf("ERROR: Cannot dump JSON string '%s' to file '%s'.\n", string(toFile), filePath)
			}
		}
	}
}

// ResotreRepoToJSONFIle() restore data form JSON file with filePath to jr
func RestoreRepoFromJSONFile(jr *jsonRepo, filePath string) {
	if filePath != "" {
		repoFile, err := os.OpenFile(filePath, os.O_RDONLY, 0777)
		if err != nil {
			fmt.Printf("ERROR: Cannot open repository file '%s' for restoring.\n", filePath)
		}
		defer repoFile.Close()
		fromFile, err := io.ReadAll(repoFile)
		if err != nil {
			fmt.Printf("ERROR: Cannot read data from repository file '%s'.\n", filePath)
		}
		err = json.Unmarshal(fromFile, &jr.JSONSlice)
		if err == nil {
			fmt.Printf("INFO: Data from repository file were restored succesfully.\n")
		} else {
			fmt.Printf("ERROR: '%s'. Data from file '%s' cannot be restored.\n", err, filePath)
		}
	}
}

func URLStorageInit(filePath string) (ur Keeper) {
	if filePath != "" {
		ur := new(FileRepo)
		ur.InitRepo(filePath)
		return ur
	} else {
		ur := new(MemRepo)
		ur.InitRepo(filePath)
		return ur
	}
}
