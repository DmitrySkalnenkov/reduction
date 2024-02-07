package userrepo

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

const DefaultUserKeyLength = 16

type UserRepo struct {
	//userMap map[int]string
	keySlice []string
}

func (repo *UserRepo) InitRepo(repoPath string) {
	//repo.keySlice = make(map[int]string)
	//defaultKey := strings.Repeat("0", config.DefaultUserKeyLength)
	repo.keySlice = make([]string, 1)
	repo.keySlice[0] = "defaultKey"
}

func (repo *UserRepo) AddUserIntoRepo() (userID int, userKey string) {
	nextUserID := len(repo.keySlice)
	key, err := UserKeyGeneration(DefaultUserKeyLength)
	if err != nil {
		fmt.Printf("ERROR: Generation of user key for UserID = %d", nextUserID)
	}
	repo.keySlice = append(repo.keySlice, key)
	return nextUserID, key
}

func UserKeyGeneration(keyLength int) (string, error) {
	b := make([]byte, keyLength)
	_, err := rand.Read(b)
	if err != nil {
		return ``, err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func (repo *UserRepo) GetKeyFromRepo(userID int) (key string) {
	return repo.keySlice[userID]
}

func (repo *UserRepo) PrintRepo() {
	fmt.Println("VVVVVVVVVVVVVVVVVVVVVVVVVV")
	fmt.Println("DEBUG: UserStorage. Begin:")
	for i := range repo.keySlice {
		fmt.Println("For UserID = %d, key is %s", i, repo.keySlice[i])
	}
	fmt.Println("DEBUG: UserStorage. End.")
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^")
}
