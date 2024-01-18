package app

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"strings"
)

const ShortURLLength = 15
const (
	HostPort string = ":8080"
	HostAddr string = "localhost"
	HostURL  string = "http://" + HostAddr + HostPort + "/"
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

func TrimSlashes(slashedStr string) string {
	return strings.ReplaceAll(slashedStr, "/", "")
}

func ReductURL(url string, shortURLLength int, urlStorage map[string]string) string {
	shortURL := randomString(shortURLLength)
	for {
		_, ok := urlStorage[shortURL]
		if !ok {
			urlStorage[shortURL] = url
			return shortURL
		}
		shortURL = randomString(shortURLLength)
	}
}
