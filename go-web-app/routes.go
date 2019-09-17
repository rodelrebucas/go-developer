package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes() {
	router.GET("/", showIndexPage)
	router.GET("/article/view/:article_id", getArticle)
	router.GET("/articles", func(c *gin.Context) {
		c.JSON(http.StatusOK, getAllArticles())
	})
}
