package storage

import (
	"fmt"
)

var URLStorage map[string]string //Storage for shortened URL

//Return saved long URL from URL storage
func GetURLFromStorage(id string, urlStorage map[string]string) string {
	url, ok := urlStorage[id]
	if ok {
		fmt.Printf("DEBUG: Found shorten URL with id '%s' is URL storage.\n", id)
		return url
	} else {
		fmt.Printf("DEBUG: Shorten URL with id '%s' not found in URL storage.\n", id)
		return ""
	}
}

func PrintMap(mapa map[string]string) {
	fmt.Println("")
	fmt.Println("DEBUG: UrlStorage:")
	for k, v := range mapa {
		fmt.Println(k, "value is", v)
	}
	fmt.Println("")
}
