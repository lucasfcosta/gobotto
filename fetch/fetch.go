package fetch

import (
	"github.com/lucasfcosta/gobotto/robotsurl"
	"io/ioutil"
	"net/http"
)

func Fetch(address string) ([]byte, error) {
	robotsAddress, err := robotsurl.RobotsURL(address)
	res, err := http.Get(robotsAddress)
	robots, err := ioutil.ReadAll(res.Body)

	res.Body.Close()

	return robots, err
}
