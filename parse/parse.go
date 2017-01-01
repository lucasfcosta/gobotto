package gobotto

import (
	"strings"
)

func Parse(text string) Robots {
	robots := newRobots()

	lines := strings.Split(text, "\n")

	var lastUserAgent string

	for _, line := range lines {
		normalized := strings.ToLower(line)

		isComment := strings.HasPrefix(normalized, "#")
		isUserAgent := strings.HasPrefix(normalized, "user-agent")
		isAllow := strings.HasPrefix(normalized, "allow")
		isDisallow := strings.HasPrefix(normalized, "disallow")

		if isComment {
			comment := strings.TrimLeft(strings.Split(line, "#")[1], " ")
			robots.Comments = append(robots.Comments, comment)
		} else if isUserAgent {
			lastUserAgent = strings.Split(line, " ")[1]
			_, exists := robots.Rules[lastUserAgent]
			if !exists {
				robots.Rules[lastUserAgent] = newRules()
			}
		} else if isAllow {
			allowedPath := strings.Split(line, " ")[1]
			subpaths := strings.Split(allowedPath, "/")
			count := 0
			for _, subpath := range subpaths {
				if subpath != "" {
					count++
				}
			}

			robots.Rules[lastUserAgent].Allow[allowedPath] = count
		} else if isDisallow {
			disallowedPath := strings.Split(line, " ")[1]

			subpaths := strings.Split(disallowedPath, "/")
			count := 0
			for _, subpath := range subpaths {
				if subpath != "" {
					count++
				}
			}

			robots.Rules[lastUserAgent].Disallow[disallowedPath] = count
		}
	}

	return robots
}
