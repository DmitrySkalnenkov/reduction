package memrepo

import (
	"fmt"
)

//(i2) Cлой storage должен имплементировать интерфейс хранения, например repositories. Это понадобится вам для подмены хранилища моком в тестах и использования DI.

type MemRepo struct {
	urlMap map[string]string
}

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
