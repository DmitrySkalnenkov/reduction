// Package entity defines main entities for business logic (services), database mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import "github.com/DmitrySkalnenkov/reduction/internal/controller/userrepo"

var URLStorage Keeper
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

type Keeper interface {
	GetURLFromRepo(token string) (string, bool)
	SetURLIntoRepo(token string, value string)
	InitRepo(repoPath string)
	PrintRepo()
}
