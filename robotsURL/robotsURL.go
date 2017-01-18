package robotsurl

import (
	"net/url"
)

func RobotsURL(address string) (string, error) {
	parsedUrl, err := url.Parse(address)
	result := parsedUrl.Scheme + "://" + parsedUrl.Host + "/robots.txt"
	return result, err
}
