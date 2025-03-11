package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.NotNil(t, testStoreService.redisClient)
}

func TestInsertionAndRetrieval(t *testing.T) {
	originalURL := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	userID := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortURL := "Jsz4k57oAX"

	SaveURLMapping(shortURL, originalURL, userID)
	retrievedURL := RetrieveOriginalURL(shortURL)

	assert.Equal(t, originalURL, retrievedURL)
}
