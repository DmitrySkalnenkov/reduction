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

type Repository struct {
	urlMap map[string]string
}

type Keeper interface {
	GetURLFromStorage(token string) (string, bool)
	SetURLIntoStorage(token string, value string)
	InitMap()
	PrintMap()
}

//var URLStorage map[string]string //Storage for shortened URL

var URLStorage Repository

// Return saved long URL from URL storage
func (repo *Repository) GetURLFromStorage(token string) (string, bool) {
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
func (repo *Repository) SetURLIntoStorage(token string, value string) {
	repo.urlMap[token] = value
}

// Initialization of Repository object
func (repo *Repository) Init() {
	repo.urlMap = make(map[string]string)
}

func (repo *Repository) PrintMap() {
	fmt.Println("VVVVVVVVVVVVVVVVVVVVVVVVVV")
	fmt.Println("DEBUG: UrlStorage. Begin:")
	for k, v := range repo.urlMap {
		fmt.Println(k, "value is", v)
	}
	fmt.Println("DEBUG: UrlStorage. End.")
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^")
}

func (repo *Repository) DumpRepositoryToJSONFile(filePath string) {
	if filePath != "" {
		repoFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
		if err != nil {
			fmt.Printf("ERROR: Cannot open repository file '%s' for dumping .\n", filePath)
		}
		defer repoFile.Close()
		if len(repo.urlMap) > 0 {
			toFile, err := json.Marshal(repo.urlMap)
			if err != nil {
				fmt.Printf("ERROR: Cannot marshal repo.urlMap '%v' to JSON.\n", repo.urlMap)
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

func (repo *Repository) RestoreRepositoryFromJSONFile(filePath string) {
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
		var tmpRepo Repository
		err = json.Unmarshal(fromFile, &tmpRepo.urlMap)
		if err == nil {
			fmt.Printf("INFO: Data from repository file were restored succesfully.\n")
			*repo = tmpRepo
		} else {
			fmt.Printf("ERROR: '%s'. Data from file '%s' cannot be restored.\n", err, filePath)
		}
	}
}
