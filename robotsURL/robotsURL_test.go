package robotsurl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRobotsURLSuccessful(t *testing.T) {
	expectedURL := "http://my-cool-domain.com/robots.txt"
	result, err := RobotsURL("http://my-cool-domain.com/blog-post/1")

	assert.Nil(t, err)
	assert.Equal(t, expectedURL, result)
}
