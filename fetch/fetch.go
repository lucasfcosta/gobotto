package gobotto

import (
	"io/ioutil"
	"net/http"
)

func Fetch(address string) ([]byte, error) {
	res, err := http.Get(address)
	robots, err := ioutil.ReadAll(res.Body)

	res.Body.Close()

	return robots, err
}
