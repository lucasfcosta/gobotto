package gobotto

import (
	robotsURL "github.com/lucasfcosta/gobotto/robotsURL"
	"io/ioutil"
	"net/http"
)

func Fetch(address string) ([]byte, error) {
	robotsAddress, err := robotsURL.RobotsURL(address)
	res, err := http.Get(robotsAddress)
	robots, err := ioutil.ReadAll(res.Body)

	res.Body.Close()

	return robots, err
}
