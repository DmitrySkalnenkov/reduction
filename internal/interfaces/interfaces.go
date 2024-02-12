package interfaces

import (
	"github.com/DmitrySkalnenkov/reduction/internal/models"
	"net/http"
)

type IDataRepo interface {
	GetURLFromRepo(token string) (string, bool)
	SetURLIntoRepo(token string, value string)
	InitRepo(repoPath string)
	PrintRepo()
}

type IDataServices interface {
	ReduceURL(urluser models.URLUser, shortURLLength int) string
	GetOriginURL(token string) (originURL string)
}

type IContentReader interface {
	GetRequestBody(w http.ResponseWriter, r *http.Request) (bodyStr string, respStatus int)
}

type IResponseWriter interface {
	SendResponse(w http.ResponseWriter, r *http.Request)
}
