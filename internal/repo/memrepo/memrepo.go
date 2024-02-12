package memrepo

import (
	"fmt"
	"github.com/DmitrySkalnenkov/reduction/internal/models"
)

//(i2) Cлой storage должен имплементировать интерфейс хранения, например repositories. Это понадобится вам для подмены хранилища моком в тестах и использования DI.

type MemRepo struct {
	urlMap map[string]models.URLUser
}

// GetURLFromRepo returns saved long URL with UserID from URL storage, and ok
func (repo *MemRepo) GetURLFromRepo(token string) (string, bool) {
	defaultURLUser := models.URLUser{URL: "", UserID: 0}
	urluser, ok := repo.urlMap[token]
	if ok {
		fmt.Printf("DEBUG: Found long URL with token '%s' is URL storage.\n", token)
		return urluser.URL, ok
	} else {
		fmt.Printf("DEBUG: Shortened URL with token '%s' not found in URL storage.\n", token)
		return defaultURLUser.URL, ok
	}
}

// SetURLIntoRepo sets url (long URL) into repository for this token(shortened URL)
func (repo *MemRepo) SetURLIntoRepo(token string, value string) {
	tmp := repo.urlMap[token]
	tmp.URL = value
	tmp.UserID = 0 //default UserID = 0
	repo.urlMap[token] = tmp
}

// InitRepo() init MemRepo object
func (repo *MemRepo) InitRepo(repoPath string) {
	repo.urlMap = make(map[string]models.URLUser)
}

func (repo *MemRepo) PrintRepo() {
	fmt.Println("VVVVVVVVVVVVVVVVVVVVVVVVVV")
	fmt.Println("DEBUG: UrlStorage. Begin:")
	for k, v := range repo.urlMap {
		fmt.Println(k, "url is", v)
	}
	fmt.Println("DEBUG: UrlStorage. End.")
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^")
}
