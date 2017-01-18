package isallowed

import (
	"github.com/lucasfcosta/gobotto/models"
	"net/url"
	"strings"
)

func IsAllowed(agent string, address string, rules models.Robots) (bool, error) {
	agentRules, specificRulesExist := rules.Rules[agent]

	if !specificRulesExist {
		var genericRulesExist bool
		agentRules, genericRulesExist = rules.Rules["*"]
		if !genericRulesExist {
			return true, nil
		}
	}

	parsedUrl, err := url.Parse(address)
	path := strings.Trim(parsedUrl.Path, "/")
	subpaths := strings.Split(path, "/")

	var lastTry string = "/"
	var allowPrecision int = 0
	var disallowPrecision int = 0
	for _, subpath := range subpaths {
		lastTry = lastTry + subpath + "/"

		foundAllowPrecision := agentRules.Allow[lastTry]
		foundDisallowPrecision := agentRules.Disallow[lastTry]

		if foundAllowPrecision > allowPrecision {
			allowPrecision = foundAllowPrecision
		}

		if foundDisallowPrecision > disallowPrecision {
			disallowPrecision = foundDisallowPrecision
		}
	}

	return allowPrecision > disallowPrecision, err
}
