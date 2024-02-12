// Package models defines main entities for business logic (services), database mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package models

import (
	"github.com/DmitrySkalnenkov/reduction/internal/repos/userrepo"
)

// var URLStorage interfaces.IDataRepo
var HostSocketAddrStr string
var BaseURLStr string
var RepoFilePathStr string
var UserKeyStorage *userrepo.UserRepo

type TxJSONMessage struct {
	URL string `json:"url"`
}

type RxJSONMessage struct {
	Result string `json:"result"`
}

type JSONLine struct {
	Token  string `json:"short_url"`
	URL    string `json:"original_url"`
	UserID string `json:"user_id"`
}

type JSONRepo struct {
	JSONSlice []JSONLine
}

type URLUser struct {
	URL    string
	UserID int
}
