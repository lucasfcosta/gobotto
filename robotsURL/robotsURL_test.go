package robotsurl

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Gets a robots.txt path for root urls
func ExampleRobotsUrl_rootURL() {
	url, _ := RobotsURL("https://example.com/")
	fmt.Println(url)
	// Output: https://example.com/robots.txt
}

// Gets a robots.txt path for any urls
func ExampleRobotsURL_anyUrl() {
	url, _ := RobotsURL("https://example.com/path/to/whatever")
	fmt.Println(url)
	// Output: https://example.com/robots.txt
}

func TestRobotsURLSuccessful(t *testing.T) {
	expectedURL := "http://my-cool-domain.com/robots.txt"
	result, err := RobotsURL("http://my-cool-domain.com/blog-post/1")

	assert.Nil(t, err)
	assert.Equal(t, expectedURL, result)
}
