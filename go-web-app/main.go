package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// create the router
	router = gin.Default()

	// load the html templates
	router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	initializeRoutes()

	// Start app
	router.Run()
}
