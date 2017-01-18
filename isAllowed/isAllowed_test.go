package isallowed

import (
	"encoding/json"
	"github.com/lucasfcosta/gobotto/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"path/filepath"
	"testing"
)

type IsAllowedTestSuite struct {
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

func (suite *IsAllowedTestSuite) TestAllowed() {
	rules := fakeRules()
	allowed, err := IsAllowed("Googlebot", "https://example.org/browse/random", rules)
	assert.True(suite.T(), allowed)
	assert.Nil(suite.T(), err)
}

func (suite *IsAllowedTestSuite) TestDisallowed() {
	rules := fakeRules()
	allowed, err := IsAllowed("Slurp", "https://example.org/thing/random", rules)
	assert.False(suite.T(), allowed)
	assert.Nil(suite.T(), err)
}

func (suite *IsAllowedTestSuite) TestGenericAllowed() {
	rules := fakeRules()
	allowed, err := IsAllowed("HipsterAgent", "https://example.org/random/whatever", rules)
	assert.True(suite.T(), allowed)
	assert.Nil(suite.T(), err)
}

func (suite *IsAllowedTestSuite) TestGenericDisallowed() {
	rules := fakeRules()
	allowed, err := IsAllowed("HipsterAgent", "https://example.org/path/whatever", rules)
	assert.False(suite.T(), allowed)
	assert.Nil(suite.T(), err)
}

func (suite *IsAllowedTestSuite) TestOverride() {
	rules := fakeRules()
	allowed, err := IsAllowed("HipsterAgent", "https://example.org/path/one", rules)
	assert.True(suite.T(), allowed)
	assert.Nil(suite.T(), err)
}

func (suite *IsAllowedTestSuite) TestSamePrecision() {
	rules := fakeRules()
	allowed, err := IsAllowed("Slurp", "https://example.org/conflicting", rules)

	// If rules conflict we disallow this URL for safety
	assert.False(suite.T(), allowed)
	assert.Nil(suite.T(), err)
}

func TestIsAllowed(t *testing.T) {
	suite.Run(t, new(IsAllowedTestSuite))
}
