package main

import (
	"fmt"

	"github.com/SonDo580/go-url-shortener/handler"
	"github.com/SonDo580/go-url-shortener/store"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to URL Shortener",
		})
	})

	r.POST("/short-url", func(c *gin.Context) {
		handler.CreateShortURL(c)
	})

	r.GET("/:shortURL", func(c *gin.Context) {
		handler.HandleShortURLRedirect(c)
	})

	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start server - Error: %v\n", err))
	}
}
