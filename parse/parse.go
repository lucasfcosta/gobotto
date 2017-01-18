package parse

import (
	"github.com/lucasfcosta/gobotto/models"
	"strings"
)

func Parse(text string) models.Robots {
	robots := models.NewRobots()

	lines := strings.Split(text, "\n")

	var lastUserAgent string

	// For each line
	for _, line := range lines {
		normalized := strings.ToLower(line)

		// Detect the semantic value of a line
		isComment := strings.HasPrefix(normalized, "#")
		isUserAgent := strings.HasPrefix(normalized, "user-agent")
		isAllow := strings.HasPrefix(normalized, "allow")
		isDisallow := strings.HasPrefix(normalized, "disallow")

		// Handle that line according to its semantic value
		if isComment {
			comment := strings.TrimLeft(strings.Split(line, "#")[1], " ")
			robots.Comments = append(robots.Comments, comment)
		} else if isUserAgent {
			lastUserAgent = strings.Split(line, " ")[1]
			_, exists := robots.Rules[lastUserAgent]
			if !exists {
				robots.Rules[lastUserAgent] = models.NewRules()
			}
		} else if isAllow {
			path := strings.Split(line, " ")[1]
			robots.Rules[lastUserAgent].Allow[path] = precision(path)
		} else if isDisallow {
			path := strings.Split(line, " ")[1]
			robots.Rules[lastUserAgent].Disallow[path] = precision(path)
		}
	}

	return robots
}

func precision(path string) int {
	subpaths := strings.Split(path, "/")
	count := 0
	for _, subpath := range subpaths {
		if subpath != "" {
			count++
		}
	}

	return count
}
