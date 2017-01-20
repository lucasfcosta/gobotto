/*
  Copy is not theft, do whatever you want with this code.
  This is an example of how we can use some extra bytes for cool reflections instead of copyright stuff.
    "The world's entire scientific and cultural heritage, published over centuries
	in books and journals, is increasingly being digitized and locked up by a handful
	of private corporations."
	-SWARTZ, Aaron
*/

/*
  Package robotsurl determines which is the URL to the robots.txt file for any site given any of its addresses.
*/
package robotsurl

import (
	"net/url"
)

func RobotsURL(address string) (string, error) {
	parsedUrl, err := url.Parse(address)
	result := parsedUrl.Scheme + "://" + parsedUrl.Host + "/robots.txt"
	return result, err
}
