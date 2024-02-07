package usecase

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"github.com/DmitrySkalnenkov/reduction/internal/entity"
	"strings"
)

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
func ReduceURL(url string, shortURLLength int, pr entity.Keeper) string {
	shortURL := randomString(shortURLLength)
	for {
		_, ok := pr.GetURLFromRepo(shortURL)
		if !ok {
			pr.SetURLIntoRepo(shortURL, url)
			return shortURL
		}
		shortURL = randomString(shortURLLength)
	}
}

//(i9) Добавьте в сервис функциональность аутентификации пользователя.
