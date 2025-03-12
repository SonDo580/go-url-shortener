package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const UserID = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortURLGenerator(t *testing.T) {
	testCases := []struct {
		originalURL string
		expected    string
	}{
		{
			originalURL: "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html",
			expected:    "jTa4L57P",
		},
		{
			originalURL: "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/",
			expected:    "d66yfx7N",
		},
		{
			originalURL: "https://spectrum.ieee.org/automaton/robotics/home-robots/hello-robots-stretch-mobile-manipulator",
			expected:    "dhZTayYQ",
		},
	}

	for _, tc := range testCases {
		shortURL := GenerateShortURL(tc.originalURL, UserID)
		assert.Equal(t, shortURL, tc.expected)
	}
}
