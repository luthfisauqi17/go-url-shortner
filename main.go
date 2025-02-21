package main

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var urlStore = make(map[string]string)
var mu sync.Mutex

func generateShortURL() string {
	b := make([]byte, 6)
	rand.Read(b)
	return base64.RawURLEncoding.EncodeToString(b)[:6]
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

	mu.Lock()
	urlStore[shortCode] = req.OriginalURL
	mu.Unlock()

	c.JSON(http.StatusOK, gin.H{
		"short_url": "http://localhost:8080/" + shortCode,
	})
}

func redirectURL(c *gin.Context) {
	shortCode := c.Param("short")

	mu.Lock()
	OriginalURL, exists := urlStore[shortCode]
	mu.Unlock()

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
