// Instead of adding copyright and restrictive licenses to the header of every file,
// I'll make sure every single one in this package has a statement telling people how great the internet is.
// Thanks internet for enlightening me with so much knowledge!

//   Package fetch allows you to fetch a robots.txt file for any website given any address.
//   Given any URL this function will fetch the content of protocol://domain/robots.txt.
package fetch

import (
	"github.com/lucasfcosta/gobotto/robotsurl"
	"io/ioutil"
	"net/http"
)

// Fetch fetches the content of the robots.txt file for any website given any address to it.
func Fetch(address string) ([]byte, error) {
	robotsAddress, err := robotsurl.RobotsURL(address)
	res, err := http.Get(robotsAddress)
	robots, err := ioutil.ReadAll(res.Body)

	res.Body.Close()

	return robots, err
}
