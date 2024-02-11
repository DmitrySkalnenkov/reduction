package userrepo

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

const DefaultUserKeyLength = 16

type UserRepo struct {
	//userMap map[int]string
	KeySlice []string
}

func (repo *UserRepo) InitRepo(repoPath string) {
	//repo.keySlice = make(map[int]string)
	//defaultKey := strings.Repeat("0", config.DefaultUserKeyLength)
	repo.KeySlice = make([]string, 1)
	repo.KeySlice[0] = "defaultKey"
}

func (repo *UserRepo) AddUserIntoRepo() (userID int, userKey string) {
	nextUserID := len(repo.KeySlice)
	key, err := UserKeyGeneration(DefaultUserKeyLength)
	if err != nil {
		fmt.Printf("ERROR: Generation of user key for UserID = %d", nextUserID)
		return 0, ""
	}
	repo.KeySlice = append(repo.KeySlice, key)
	return nextUserID, key
}

func UserKeyGeneration(keyLength int) (string, error) {
	b := make([]byte, keyLength)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func (repo *UserRepo) GetKeyFromRepo(userID int) (key string) {
	if userID >= len(repo.KeySlice) || userID < 0 {
		fmt.Printf("ERROR: UserID (%d) must be in the range of indexes of UserRepo %d - %d.\n", userID, 0, len(repo.KeySlice))
		fmt.Println("INFO: Empty key will be used.")
		return ""
	}
	key = repo.KeySlice[userID]
	return key
}

func (repo *UserRepo) PrintRepo() {
	fmt.Println("VVVVVVVVVVVVVVVVVVVVVVVVVV")
	fmt.Println("DEBUG: UserStorage. Begin:")
	for i := range repo.KeySlice {
		fmt.Printf("For UserID = %d, key is %s.\n", i, repo.KeySlice[i])
	}
	fmt.Println("DEBUG: UserStorage. End.")
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^")
}
