package cookies

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/DmitrySkalnenkov/reduction/internal/entity"
	"net/http"
	"strconv"
	"strings"
)

func GenerateSign(userIDStr string, key string) (sign string) {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(userIDStr))
	signBytes := h.Sum(nil)
	sign = hex.EncodeToString(signBytes)
	return sign
}

func GenerateSignedCookie(userID int, key string) (c *http.Cookie) {
	var cookie http.Cookie
	cookie = http.Cookie{
		Name: "auth_cookie",
	}
	userIDStr := fmt.Sprintf("%08d", userID)
	signStr := GenerateSign(userIDStr, key)
	cookie.Value = fmt.Sprintf("%s-%s", userIDStr, signStr)
	return &cookie
}

func isAuthSignCorrect(userID int, receivedSign string) bool {
	key := entity.UserKeyStorage.GetKeyFromRepo(userID)
	userIDStr := fmt.Sprintf("%08d", userID)
	expectedSign := GenerateSign(userIDStr, key)
	if receivedSign == expectedSign {
		return true
	} else {
		return false
	}
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

/*func GetAuthCookie(r *http.Request) (cookie *http.Cookie, err error) {
	cookie, err = r.Cookie("auth_cookie")
	return cookie, err
}*/

func GetAuthUserID(w http.ResponseWriter, r *http.Request) int {
	var curUserID int
	var curUserKey string
	//Get auth cookie
	authCookie, err := r.Cookie("auth_cookie")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			fmt.Printf("INFO: Auth cookie is not found. New auth cookie will be generated.\n")
			curUserID, curUserKey = entity.UserKeyStorage.AddUserIntoRepo()
			authCookie = GenerateSignedCookie(curUserID, curUserKey)
			http.SetCookie(w, authCookie) //Set cookie for new user
			return curUserID
		} else {
			fmt.Printf("ERROR: Error to get cookie. Err - %v.\n", err)
			w.WriteHeader(http.StatusBadRequest)
			curUserID = 0 //Default UserID
			return curUserID
		}
	}
	curUserID, signStr := ParseAuthCookie(authCookie)
	if isAuthSignCorrect(curUserID, signStr) {
		return curUserID
	} else {
		fmt.Printf("INFO: Auth cookie is not found. New auth cookie will be generated.\n")
		curUserID, curUserKey = entity.UserKeyStorage.AddUserIntoRepo()
		authCookie = GenerateSignedCookie(curUserID, curUserKey)
		http.SetCookie(w, authCookie) //Set cookie for new user
		return curUserID
	}
}
