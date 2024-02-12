package interfaces

type DataRepo interface {
	GetURLFromRepo(token string) (string, bool)
	SetURLIntoRepo(token string, value string)
	InitRepo(repoPath string)
	PrintRepo()
}
