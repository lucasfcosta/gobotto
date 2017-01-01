/*
Package gobotto offers a simple library for fetching, parsing and dealing with `robots.txt` files.

Gobotto has the following methods:

	* gobotto.RobotsURL(URL string)
	* gobotto.Fetch(URL string)
	* gobotto.Parse(robotsContents string)
	* gobotto.IsAllowed(userAgent string, route string, rules RobotsRules)
*/
package gobotto
