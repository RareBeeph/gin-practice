package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	// Minor nitpick - HTTP paths should be as lowercase as possible
	// They ARE case sensitive and nobody remembers the casing correctly
	// This becomes far more important the larger your thing is
	router.GET("/GetEndpoint", func(c *gin.Context) {
		c.String(http.StatusOK, "Placeholder response (Get)")
	})

	router.POST("/PostEndpoint", func(c *gin.Context) {
		postData, _ := io.ReadAll(c.Request.Body)
		// There's a number of ways to do this particular task, but this is
		// the most flexible in my opinion and doesn't lock you into values
		// that can be parsed as strings
		c.Data(http.StatusOK, c.Request.Header.Get("content-type"), postData)
	})

	router.DELETE("/DeleteEndpoint", func(c *gin.Context) {
		c.String(http.StatusOK, "Placeholder response (Delete)")
	})

	return router
}

func main() {
	router := setupRouter()
	router.Run()
}
