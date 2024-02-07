package cookies

import (
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/DmitrySkalnenkov/reduction/internal/entity"
	"net/http"
	"strconv"
	"strings"
)

func GenerateSignedCookie(userID int, key string) (c *http.Cookie) {
	var cookie http.Cookie
	cookie = http.Cookie{
		Name: "auth_cookie",
	}
	userIDStr := fmt.Sprintf("%08d", userID)
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(userIDStr))
	sign := h.Sum(nil)
	cookie.Value = fmt.Sprintf("%s-%s", userIDStr, sign)
	return &cookie
}

func ParseAuthCookie(cookie *http.Cookie) (userID int, sign string) {
	var err error
	valueSlice := strings.Split(cookie.Value, "-")
	userID, err = strconv.Atoi(valueSlice[0])
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		userID = 0 //Default userID
		sign = ""  //Default sign
	}
	sign = valueSlice[1]
	return userID, sign
}

func GetAuthCookie(r *http.Request) (cookie *http.Cookie, err error) {
	cookie, err = r.Cookie("auth_cookie")
	return cookie, err
}

func GetAuthUserID(w http.ResponseWriter, r *http.Request) (curUserID int) {
	var userKey string
	//Get auth cookie
	authCookie, err := GetAuthCookie(r)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			fmt.Printf("INFO: Auth cookie is not found. New auth cookie will be generated.\n")
			curUserID, userKey = entity.UserKeyStorage.AddUserIntoRepo()
			authCookie = GenerateSignedCookie(curUserID, userKey)
		} else {
			fmt.Printf("ERROR: Error to get cookie. Err - %v.\n", err)
			w.WriteHeader(http.StatusBadRequest)
			curUserID = 0 //Default UserID
			return curUserID
		}
	}
	curUserID, _ = ParseAuthCookie(authCookie)
	//userKey = entity.UserKeyStorage.GetKeyFromRepo(userID)
	//TODO: Check cookie sign
	return curUserID
}
