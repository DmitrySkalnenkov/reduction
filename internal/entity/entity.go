// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

type TxJSONMessage struct {
	URL string `json:"url"`
}
type RxJSONMessage struct {
	Result string `json:"result"`
}

type JSONLine struct {
	Token string `json:"token"`
	URL   string `json:"url"`
}

type JSONRepo struct {
	JSONSlice []JSONLine
}

type Keeper interface {
	GetURLFromRepo(token string) (string, bool)
	SetURLIntoRepo(token string, value string)
	InitRepo(repoPath string)
	PrintRepo()
}
