package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowIndexPageUnauthenticated(t *testing.T) {
	router := getRouter(true)

	// Create a router for /
	router.GET("/", showIndexPage)
	// Create a get request to  /
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
		// Test  http status code == 200
		statusOK := w.Code == http.StatusOK

		// Test if title == Home Page
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Article List</title>") > 0

		return statusOK && pageOK
	})
}

func TestSingleArticleIsDisplayed(t *testing.T) {
	articleInst, error := getArticleByID(3)
	if error != nil {
		t.Fail()
	}
	// setup router and request
	router := getRouter(true)
	router.GET("/article/view/:article_id", getArticle)
	req, _ := http.NewRequest("GET", "/article/view/3", nil)

	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
		// Test  http status code == 200
		statusOK := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), articleInst.Content) > 0
		return statusOK && pageOK
	})

}
