package userrepo

import (
	"fmt"
	"testing"
)

func TestUserRepo_AddUserIntoRepo(t *testing.T) {
	var ur UserRepo
	ur.InitRepo("")
	ur = UserRepo{
		KeySlice: []string{
			"123456789012345",
		},
	}

	tests := []struct {
		name       string
		wantUserID int
	}{ //Test table
		{
			name:       "Positive test. UserID = 1 ",
			wantUserID: 1,
		},
		{
			name:       "Positive test. UserID = 2 ",
			wantUserID: 2,
		},
	}

	for _, tt := range tests {
		var resultUserID int
		var resultUserKey string
		t.Run(tt.name, func(t *testing.T) {
			resultUserID, resultUserKey = ur.AddUserIntoRepo()
			fmt.Printf("TEST_DEBUG: New userID is  '%d' with key  '%s'.\n", resultUserID, resultUserKey)
			if resultUserID != tt.wantUserID && len(resultUserKey) != DefaultUserKeyLength {
				t.Errorf("TEST_ERROR: Result userID is %d and lenght is %d, but wanted %d and %d.\n", resultUserID, len(resultUserKey), tt.wantUserID, DefaultUserKeyLength)
			}
		})
	}

}

func TestUserRepo_GetKeyFromRepo(t *testing.T) {
	var ur UserRepo
	ur.InitRepo("")
	ur = UserRepo{
		KeySlice: []string{
			"1234567890123456",
			"6543210987654321",
			"1234567887654321",
		},
	}

	tests := []struct {
		name        string
		inputUserID int
		wantUserKey string
	}{ //Test table
		{
			name:        "Positive test. UserID = 1",
			inputUserID: 1,
			wantUserKey: "6543210987654321",
		},
		{
			name:        "Positive test. UserID = 2",
			inputUserID: 2,
			wantUserKey: "1234567887654321",
		},
		{
			name:        "Negative test. UserID = 3. No such UserID",
			inputUserID: 3,
			wantUserKey: "",
		},
	}

	for _, tt := range tests {
		var resultUserKey string
		t.Run(tt.name, func(t *testing.T) {
			resultUserKey = ur.GetKeyFromRepo(tt.inputUserID)
			fmt.Printf("TEST_DEBUG: UserKey for UserID '%d' is '%s'.\n", tt.inputUserID, resultUserKey)
			if resultUserKey != tt.wantUserKey && len(resultUserKey) != DefaultUserKeyLength {
				t.Errorf("TEST_ERROR: UserKey '%s' for UserID '%d' but wanted %s.\n", resultUserKey, tt.inputUserID, tt.wantUserKey)
			}
		})
	}
}
