package parse

import (
	"encoding/json"
	"fmt"
	"github.com/lucasfcosta/gobotto/fetch"
	"github.com/lucasfcosta/gobotto/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"path/filepath"
	"testing"
)

// Parses a robots.txt file.
func ExampleParse() {
	// This fetches robots.txt from https://example.com/robots.txt
	robots, _ := fetch.Fetch("https://example.com")

	// This parses the fetched rules
	parsed := Parse(string(robots))

	fmt.Println(parsed)
}

type ParseTestSuite struct {
	suite.Suite
}

func fakeText() string {
	path, _ := filepath.Abs("../fake/fake_robots.txt")
	robots, _ := ioutil.ReadFile(path)
	return string(robots)
}

func fakeRules() models.Robots {
	path, _ := filepath.Abs("../fake/fake_robots.json")
	robotsJson, _ := ioutil.ReadFile(path)

	var rules models.Robots
	json.Unmarshal(robotsJson, &rules)
	return rules
}

func (suite *ParseTestSuite) TestParse() {
	result := Parse(fakeText())
	expected := fakeRules()
	assert.Equal(suite.T(), expected, result)
}

func TestParseSuite(t *testing.T) {
	suite.Run(t, new(ParseTestSuite))
}
