package handler

import (
	"net/http"

	"github.com/SonDo580/go-url-shortener/shortener"
	"github.com/SonDo580/go-url-shortener/store"
	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	LongURL string `json:"long_url" binding:"required"`
	UserID  string `json:"user_id" binding:"required"`
}

func CreateShortURL(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortURL := shortener.GenerateShortURL(creationRequest.LongURL, creationRequest.UserID)
	store.SaveURLMapping(shortURL, creationRequest.LongURL, creationRequest.UserID)

	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":   "short URL created successfully",
		"short_url": host + shortURL,
	})
}

func HandleShortURLRedirect(c *gin.Context) {
	shortURL := c.Param("shortURL")
	originalURL := store.RetrieveOriginalURL(shortURL)
	c.Redirect(302, originalURL)
}
