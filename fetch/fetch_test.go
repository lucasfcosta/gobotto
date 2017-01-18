package fetch

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

type FetchTestSuite struct {
	suite.Suite
}

func (suite *FetchTestSuite) TestFetchAnyPath() {
	// Create fake server
	responseCallback := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		assert.Equal(suite.T(), "/robots.txt", r.URL.Path)
	}
	mockServer := httptest.NewServer(http.HandlerFunc(responseCallback))
	defer mockServer.Close()

	// Fetch robots.txt from fake server
	Fetch(mockServer.URL)
}

func (suite *FetchTestSuite) TestFetchSuccessful() {
	// Read local robots.txt file
	path, _ := filepath.Abs("../fake/fake_robots.txt")
	robots, _ := ioutil.ReadFile(path)

	// Create fake server
	responseCallback := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(robots))
	}
	mockServer := httptest.NewServer(http.HandlerFunc(responseCallback))
	defer mockServer.Close()

	// Fetch robots.txt from fake server
	response, err := Fetch(mockServer.URL)

	assert.Nil(suite.T(), err, "There should be no error")
	assert.Equal(suite.T(), robots, response)
}

func TestFetchSuite(t *testing.T) {
	suite.Run(t, new(FetchTestSuite))
}
