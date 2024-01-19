package storage

import (
	"fmt"
)

//(i2) Cлой storage должен имплементировать интерфейс хранения, например repositories. Это понадобится вам для подмены хранилища моком в тестах и использования DI.

type JSONMessage struct {
	Name  string
	Value string
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
	fmt.Println("")
	fmt.Println("DEBUG: UrlStorage:")
	for k, v := range repo.urlMap {
		fmt.Println(k, "value is", v)
	}
	fmt.Println("")
}
