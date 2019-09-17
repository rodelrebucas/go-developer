package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpArticleList []article

// setup
func TestMain(m *testing.M) {

	// set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// run other tests
	os.Exit(m.Run())

}

// create router for testing
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

// process request and test the response
func testHTTPResponse(t *testing.T, router *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	w := httptest.NewRecorder()

	// run the router with recorder and request
	router.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// save list to temp
func saveLists() {
	tmpArticleList = articleList
}

// restore list
func restoreLists() {
	articleList = tmpArticleList
}
