package main

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

var urlStore = make(map[string]string)

func generateShortURL() string {
	b := make([]byte, 6)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:6]
}

func shortenURL(c *gin.Context) {
	var req struct {
		OriginalURL string `json:"original_url"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	shortCode := generateShortURL()

	urlStore[shortCode] = req.OriginalURL

	c.JSON(http.StatusOK, gin.H{
		"short_url": "http://localhost:8080/" + shortCode,
	})
}

func redirectURL(c *gin.Context) {
	shortCode := c.Param("short")

	OriginalURL, exists := urlStore[shortCode]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusFound, OriginalURL)
}

func main() {
	r := gin.Default()
	r.POST("/shorten", shortenURL)
	r.GET("/:short", redirectURL)

	r.Run(":8080")
}
