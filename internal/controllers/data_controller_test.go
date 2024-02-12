package controllers

import (
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPlainReader_GetRequestBody(t *testing.T) {
	req := httptest.NewRequest("POST", "http://local.com", strings.NewReader("http://yandex.com/urlforreduction"))
	req.Header.Set("Content-Type", "plain/text")
	pr := &plainReader{}
	w := httptest.NewRecorder()
	curBodyStr, curHTTPStatus := pr.GetRequestBody(w, req)
	fmt.Printf("TEST_INFO: Current body string - %s, current HTTP-status - %d.\n", curBodyStr, curHTTPStatus)
}

func TestJSONReader_GetRequestBody(t *testing.T) {
	req := httptest.NewRequest("POST", "http://local.com", strings.NewReader("{\"url\":\"http://yandex.com/urlforreduction\"}"))
	req.Header.Set("Content-Type", "application/json")
	pr := &jsonReader{}
	w := httptest.NewRecorder()
	curBodyStr, curHTTPStatus := pr.GetRequestBody(w, req)
	fmt.Printf("TEST_INFO: Current body string - %s, current HTTP-status - %d.\n", curBodyStr, curHTTPStatus)
}

func TestContentReaderInit(t *testing.T) {
	req := httptest.NewRequest("POST", "http://local.com", strings.NewReader("{\"url\":\"http://yandex.com/urlforreduction\"}"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	cr := contentReaderInit(req)
	curBodyStr, curHTTPStatus := cr.GetRequestBody(w, req)
	fmt.Printf("TEST_INFO: Current body string - %s, current HTTP-status - %d.\n", curBodyStr, curHTTPStatus)
}
