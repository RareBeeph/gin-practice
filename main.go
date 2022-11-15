package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/GetEndpoint", func(c *gin.Context) {
		c.String(http.StatusOK, "Placeholder response (Get)")
	})

	router.POST("/PostEndpoint", func(c *gin.Context) {
		postData, _ := io.ReadAll(c.Request.Body)
		c.String(http.StatusOK, (string)(postData)) // Placeholder response: return raw data received
	})

	router.DELETE("/DeleteEndpoint", func(c *gin.Context) {
		c.String(http.StatusOK, "Placeholder response (Delete)")
	})

	router.Run()
}
