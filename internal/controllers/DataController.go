package controllers

import (
	"github.com/DmitrySkalnenkov/reduction/internal/interfaces"
	"github.com/DmitrySkalnenkov/reduction/internal/models"
	"github.com/DmitrySkalnenkov/reduction/internal/services"
)

type DataController struct {
	UserID         int
	ShortURLLength int
	DataRepo       interfaces.DataRepo
}

func (dc *DataController) GetShortURL(url string) (token string) {
	urluser := models.URLUser{URL: url, UserID: dc.UserID}
	token = services.ReduceURL(urluser, dc.ShortURLLength, dc.DataRepo)
	return token
}

func (dc *DataController) GetOriginURL(token string) (originURL string) {
	originURL = services.GetOriginURL(token, dc.DataRepo)
	return originURL
}
