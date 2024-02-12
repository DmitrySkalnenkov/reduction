package controllers

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/DmitrySkalnenkov/reduction/config"
	"github.com/DmitrySkalnenkov/reduction/internal/repos"

	//"github.com/DmitrySkalnenkov/reduction/internal/controllers/http/cookies"
	"github.com/DmitrySkalnenkov/reduction/internal/interfaces"
	"github.com/DmitrySkalnenkov/reduction/internal/models"
	"github.com/DmitrySkalnenkov/reduction/internal/repos/filerepo"
	"github.com/DmitrySkalnenkov/reduction/internal/repos/memrepo"
	"github.com/DmitrySkalnenkov/reduction/internal/repos/userrepo"
	"github.com/DmitrySkalnenkov/reduction/internal/services"
	"io"
	"log"
	"net/http"
	"strings"
)

type DataController struct {
	contentType string
	interfaces.IDataServices
}

func InitDataController() DataController {
	var URLStorage repos.DataRepo
	if models.RepoFilePathStr == "" {
		URLStorage := new(memrepo.MemRepo)
		URLStorage.InitRepo("")
	} else {
		URLStorage := new(filerepo.FileRepo)
		URLStorage.InitRepo(models.RepoFilePathStr)
	}
	models.UserKeyStorage = new(userrepo.UserRepo) //For user keys
	models.UserKeyStorage.InitRepo("")

	dataService := &services.DataService{IDataRepo: URLStorage}
	dataController := DataController{IDataServices: dataService}
	return dataController
}

type plainReader struct {
	reader io.Reader
}

type jsonReader struct {
	decoder *json.Decoder
}

type zipContentReader struct {
	reader io.Reader
}

func contentReaderInit(r *http.Request) (cr interfaces.IContentReader, contentTypeResp string) {
	contentTypeStr := r.Header.Get("Content-Type")
	if strings.Contains(contentTypeStr, "application/json") {
		cr = &jsonReader{}
		contentTypeResp = "json"
	} else if strings.Contains(contentTypeStr, "zip") {
		cr = &zipContentReader{}
		contentTypeResp  = "gzip"
	} else {
		cr = &plainReader{}
		contentTypeResp = "plain"
	}
	return cr, contentTypeResp
}

func contentWriterInit(w http.ResponseWriter, contentType string, data string, respStatus int) (cw interfaces.IResponseWriter){
	switch contentType {
		case "json": {
			cw = &jsonWriter{}

		}
		case "gzip": {
			cw = &gzipWriter{}
		}
		default:	{
			cw = &plainWriter{}
		}
	}
}

func (c *plainReader) GetRequestBody(w http.ResponseWriter, r *http.Request) (bodyStr string, respStatus int) {
	c.reader = r.Body
	body, err := io.ReadAll(c.reader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return "", http.StatusBadRequest
	}
	bodyStr = string(body)
	fmt.Printf("DEBUG: POST request body is: '%s'\n", bodyStr)
	respStatus = http.StatusCreated
	return bodyStr, respStatus
}

func (c *jsonReader) GetRequestBody(w http.ResponseWriter, r *http.Request) (bodyStr string, respStatus int) {
	var curJSONMsg models.TxJSONMessage
	c.decoder = json.NewDecoder(r.Body)
	err := c.decoder.Decode(&curJSONMsg)
	if err != nil {
		log.Printf("ERROR: JSON decoding occured, %s.\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return "", http.StatusBadRequest
	} else if curJSONMsg.URL != "" {
		bodyStr = curJSONMsg.URL
		fmt.Printf("DEBUG: JSON body: URL = '%s'.\n", bodyStr)
		respStatus = http.StatusCreated
	} else {
		bodyStr = curJSONMsg.URL
		respStatus = http.StatusNotFound
	}
	return bodyStr, respStatus
}

func (c *zipContentReader) GetRequestBody(w http.ResponseWriter, r *http.Request) (bodyStr string, respStatus int) {
	gz, err := gzip.NewReader(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Printf("ERROR: %v", err)
		return "", http.StatusInternalServerError
	}
	c.reader = gz
	defer gz.Close()
	return bodyStr, respStatus
}

type plainWriter struct {
	response string
	respStatus int
	writer io.Writer
}

func (c *plainWriter) SendResponse(w http.ResponseWriter, r *http.Request,  curRespStatus int) (bodyStr string, respStatus int) {
	w.WriteHeader(curRespStatus)
	_, err := w.Write([]byte(models.BaseURLStr + "/" + c.response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (dc *DataController) ReduceURLAction(w http.ResponseWriter, r *http.Request) {
	var err error
	cr, contentType  := contentReaderInit(r)

	curBodyStr, curRespStatus := cr.GetRequestBody(w, r)
	curURLUser := models.URLUser{URL: curBodyStr, UserID: 0}

	//curUserID := cookies.GetAuthUserID(w, r)
	//curURLUser := models.URLUser{URL: curBodyStr, UserID: curUserID}
	token := dc.IDataServices.ReduceURL(curURLUser, config.DefaultShortURLLength)
	cw := contentWriterInit(w, dc.contentType, token, curRespStatus)
	cw.SendResponse(w,r)
	//fmt.Printf("DEBUG: Shortened URL for UserID = %d is: '%s'.\n", curUserID, token)
	w.WriteHeader(curRespStatus)
	_, err = w.Write([]byte(models.BaseURLStr + "/" + token))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

/*
type DataController struct {
	UserID         int
	ShortURLLength int
	DataRepo       interfaces.IDataRepo
}

func (dc *DataController) GetShortURL(url string) (token string) {
	urluser := models.URLUser{URL: url, UserID: dc.UserID}
	token = services.ReduceURL(urluser, dc.ShortURLLength, dc.DataRepo)
	services.
	return token
}

func (dc *DataController) GetOriginURL(token string) (originURL string) {
	originURL = services.GetOriginURL(token, dc.DataRepo)
	return originURL
}
*/
