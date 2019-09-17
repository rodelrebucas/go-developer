package main

import "testing"

func isLengthEqual(aList []article) bool {
	return len(aList) == len(articleList)
}

func isContentIdentical(aList []article) bool {
	identical := true
	for i, v := range aList {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {
			identical = false
		}
	}
	return identical
}
func TestGetAllArticles(t *testing.T) {
	aList := getAllArticles()

	if !isLengthEqual(aList) {
		t.Fail()
	}

	if !isContentIdentical(aList) {
		t.Fail()
	}

}
