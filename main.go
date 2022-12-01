package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/getendpoint", func(c *gin.Context) {
		c.String(http.StatusOK, "Placeholder response (Get)")
	})

	router.POST("/postendpoint", func(c *gin.Context) {
		postData, _ := io.ReadAll(c.Request.Body)
		c.Data(http.StatusOK, c.Request.Header.Get("content-type"), postData) // Placeholder response: return raw data received
	})

	router.DELETE("/deleteendpoint", func(c *gin.Context) {
		c.String(http.StatusOK, "Placeholder response (Delete)")
	})

	return router
}

func main() {
	router := setupRouter()
	router.Run()
}
