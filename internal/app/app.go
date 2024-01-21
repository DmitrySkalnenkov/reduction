package app

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"os"
	"strings"

	"github.com/DmitrySkalnenkov/reduction/internal/storage"
)

const ShortURLLength = 15
const (
	DefaultHostPort string = ":8080"
	DefaultHostAddr string = "localhost"
	HostURL         string = "http://" + DefaultHostAddr + DefaultHostPort + "/"
)

var HostPortStr string
var BaseURLStr string

// (i5) Добавьте возможность конфигурировать сервис с помощью переменных окружения:
// Get enviroment variables SERVER_ADDRESS and BASE_URL check values and set them to global variables HostPortStr and BaseURLStr
func GetEnv() {
	envHostAddr, isEnvHostAddr := os.LookupEnv("SERVER_ADDRESS") //(i5) адрес запуска HTTP-сервера с помощью переменной SERVER_ADDRESS (видимо адрес и порт)
	if isEnvHostAddr && envHostAddr != "" {
		HostPortStr = envHostAddr //+ DefaultHostPort
		fmt.Printf("DEBUG: Found SERVER_ADDRESS enviroment variable '%s', HostPortStr = '%s'\n", envHostAddr, HostPortStr)
	} else {
		HostPortStr = DefaultHostAddr + DefaultHostPort //(i1) Сервер должен быть доступен по адресу: http://localhost:8080.
		fmt.Printf("DEBUG: Will be used default host and port, HostPortStr = '%s'\n", HostPortStr)
	}
	envBaseURL, isEnvBaseURL := os.LookupEnv("BASE_URL") //(i5) базовый адрес результирующего сокращённого URL с помощью переменной BASE_URL.
	if isEnvBaseURL && envBaseURL != "" {
		BaseURLStr = envBaseURL + "/"
		fmt.Printf("DEBUG: Found BASE_URL enviroment variable  '%s', BaseURLStr = '%s'\n", envBaseURL, BaseURLStr)
	} else {
		BaseURLStr = "http://" + HostPortStr + "/"
		fmt.Printf("DEBUG: Will be used default BaseURLStr = '%s'\n", BaseURLStr)
	}
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

func TrimSlashes(slashedStr string) string {
	return strings.ReplaceAll(slashedStr, "/", "")
}

// Reduce long URL to token, save token as key and URL as value into URLStorage, return token.
func ReduceURL(url string, shortURLLength int, pr *storage.Repository) string {
	shortURL := randomString(shortURLLength)
	for {
		_, ok := pr.GetURLFromStorage(shortURL)
		if !ok {
			//urlStorage[shortURL] = url
			pr.SetURLIntoStorage(shortURL, url)
			return shortURL
		}
		shortURL = randomString(shortURLLength)
	}
}
