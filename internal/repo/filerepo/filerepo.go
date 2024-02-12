package filerepo

import (
	"encoding/json"
	"fmt"
	"github.com/DmitrySkalnenkov/reduction/internal/models"
	"io"
	"os"
	"strconv"
)

type FileRepo struct {
	urlFile     *os.File
	urlFilePath string
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
		//var jsonData models.JSONRepo
		var jsonData []models.JSONLine
		err = json.Unmarshal(dataFromFile, &jsonData)
		if err != nil {
			fmt.Printf("ERROR: Filed unmarshal data from repo file.\n")
			return
		} else {
			fmt.Printf("==repo-begin=============%s=================\n", repo.urlFilePath)
			for i := 0; i < len(jsonData); i++ {
				fmt.Printf("'%s':'%s'\n", jsonData[i].Token, jsonData[i].URL)
			}
			fmt.Printf("==repo-end===============%s=================\n", repo.urlFilePath)
		}
	} else {
		fmt.Printf("ERROR: FileRepo path is empty. Cannot be used filerepository.\n")
		return
	}
}

// SetURLIntoRepo() Save token and url in JSON foramat into JSON file
func (repo *FileRepo) SetURLIntoRepo(token string, value string) {
	if repo.urlFilePath != "" {
		var tmpJSONRepo models.JSONRepo
		RestoreRepoFromJSONFile(&tmpJSONRepo, repo.urlFilePath)
		curJSONLine := models.JSONLine{Token: token, URL: value, UserID: fmt.Sprintf("%d", 0)} //UserID = 0 - default UserID
		for i := 0; i < len(tmpJSONRepo.JSONSlice); i++ {
			if curJSONLine.Token == tmpJSONRepo.JSONSlice[i].Token {
				fmt.Printf("INFO: Token with such urluser ('%s') already in JSON file repository. Token should be unique. The memrepo didn'token change.\n", curJSONLine.Token)
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
	var curURLUser models.URLUser
	if repo.urlFilePath != "" {
		var tmpJSONRepo models.JSONRepo
		RestoreRepoFromJSONFile(&tmpJSONRepo, repo.urlFilePath)
		for i := 0; i < len(tmpJSONRepo.JSONSlice); i++ {
			if token == tmpJSONRepo.JSONSlice[i].Token {
				fmt.Printf("INFO: Token ('%s') was found in JSON file repository\n", token)
				isURLExists = true
				curUserID, err := strconv.Atoi(tmpJSONRepo.JSONSlice[i].UserID)
				if err != nil {
					fmt.Printf("ERROR: conver string UserID %s to int.\n", tmpJSONRepo.JSONSlice[i].UserID)
				}
				curURLUser = models.URLUser{URL: tmpJSONRepo.JSONSlice[i].URL, UserID: curUserID}
				return curURLUser.URL, isURLExists
			}
		}
	}
	return curURLUser.URL, isURLExists
}

// DumpRepoToJSONFIle() dumps data form jr to JSON file with filePath
func DumpRepoToJSONFile(jr *models.JSONRepo, filePath string) {
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
				fmt.Printf("ERROR: Cannot marshal memrepo.urlMap '%v' to JSON.\n", jr)
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
func RestoreRepoFromJSONFile(jr *models.JSONRepo, filePath string) {
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
