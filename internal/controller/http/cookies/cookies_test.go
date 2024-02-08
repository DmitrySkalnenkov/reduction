package cookies

import (
	"github.com/DmitrySkalnenkov/reduction/internal/controller/userrepo"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAuthUserID(t *testing.T) {
	var ur userrepo.UserRepo
	ur.InitRepo("")
	ur = userrepo.UserRepo{
		KeySlice: []string{
			"123456789012345",
		},
	}

	tests := []struct {
		name        string
		inputCookie *http.Cookie
		wantUserID  int
	}{ //Test table
		{
			name: "Positive test 1.",
			inputCookie: &http.Cookie{
				Name:     "auth_cookie",
				Value:    "1-qwertyuiopdfghj",
				Path:     "/tmp",
				HttpOnly: true,
				Secure:   true,
			},
			wantUserID: 1,
		},
		{
			name: "Positive test 2.",
			inputCookie: &http.Cookie{
				Name:     "auth_cookie",
				Value:    "1-qwertyuklhouyioyyiopdfghj",
				Path:     "/tmp",
				HttpOnly: true,
				Secure:   true,
			},
			wantUserID: 1,
		},
	}

	for _, tt := range tests {
		// запускаем каждый тест
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "http://localhost:8080", nil)
			w := httptest.NewRecorder()
			req.AddCookie(tt.inputCookie)
			resultUserID := GetAuthUserID(w, req)
			if resultUserID != tt.wantUserID {
				t.Errorf("TEST_ERROR: Result UserID (%d) and wanted UserID (%d) is not equal.\n", resultUserID, tt.wantUserID)
			}
		})
	}
}

func TestGenerateSignedCookie(t *testing.T) {
	type inputs struct {
		userID int
		key    string
	}

	tests := []struct {
		name         string
		inputUserID  int
		inputUserKey string
		wantedCookie *http.Cookie
	}{ //Test table
		{
			name:         "Positive test 1",
			inputUserID:  1,
			inputUserKey: "qwerty",
			wantedCookie: &http.Cookie{
				Name:  "auth_cookie",
				Value: "00000001-6d18f785d34d9d608fa7743f5006604e3bbe762225e9551566940e6da5886f16",
			},
		},
		{
			name:         "Positive test 2",
			inputUserID:  99999999,
			inputUserKey: "qwerty",
			wantedCookie: &http.Cookie{
				Name:  "auth_cookie",
				Value: "99999999-7ae7584828bbe525027a172ed86ce2fcfee261293a2a6ce158bd16740da92e50",
			},
		},
	}

	for _, tt := range tests {
		// запускаем каждый тест
		t.Run(tt.name, func(t *testing.T) {
			var resultCookie *http.Cookie
			resultCookie = GenerateSignedCookie(tt.inputUserID, tt.inputUserKey)
			if resultCookie.Value != tt.wantedCookie.Value {
				t.Errorf("TEST_ERROR: Result cookie value (%s) is not expected. Wanted is (%s).\n", resultCookie.Value, tt.wantedCookie.Value)
			}
		})
	}
}
