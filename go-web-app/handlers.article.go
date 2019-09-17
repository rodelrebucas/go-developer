package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()
	render(c, gin.H{
		"title":   "Article List",
		"payload": articles,
	}, "index.html")
}

func getArticle(c *gin.Context) {
	// Check if article_id is an integer
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if article_id exist
		if article, err := getArticleByID(articleID); err == nil {
			// render to view
			render(
				c,
				gin.H{
					"Title":   "Single Article",
					"payload": article,
				},
				"article.html",
			)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
