package services

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"github.com/DmitrySkalnenkov/reduction/internal/interfaces"
	"github.com/DmitrySkalnenkov/reduction/internal/models"
	"strings"
)

type DataService struct {
	interfaces.IDataRepo
}

func randomString(length int) string {
	randomBytes := make([]byte, 32)
	if length > 0 {
		_, err := rand.Read(randomBytes)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Printf("ERROR: Random string length should be > 0. Now it is %d.\n", length)
		return ""
	}
	randFixStr := base32.StdEncoding.EncodeToString(randomBytes)[:length]
	return strings.ToLower(randFixStr)
}

// ReduceURL() reduce long URL to token, save token as key and URL as value into URLStorage, return token.
func (ds *DataService) ReduceURL(urluser models.URLUser, shortURLLength int) string {
	shortURL := randomString(shortURLLength)
	for {
		_, ok := ds.GetURLFromRepo(shortURL)
		if !ok {
			ds.SetURLIntoRepo(shortURL, urluser.URL)
			ds.PrintRepo() //For DEBUG
			return shortURL
		}
		shortURL = randomString(shortURLLength)
	}
}

// GetOriginURL() gets saved long url by token
func (ds *DataService) GetOriginURL(token string) (originURL string) {
	originURL, _ = ds.GetURLFromRepo(token)
	return originURL
}

//TODO: add userID into GetURLFromRepo
