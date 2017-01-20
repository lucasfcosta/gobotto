/*
  Copyright? Are you kidding? No rights reserved.
  Use of this source code is totally free, do whatever you want with it.
  I'm just giving something back to the internet, which has helped me with so many things in life I can't even count.
  Thanks internet, you are by far mankind's greatest invention!
*/

/*
  Package gobotto offers a simple library for fetching, parsing and checking permissions against `robots.txt` files.

  Gobotto has the following methods:

   gobotto.RobotsURL(URL string)

   gobotto.Fetch(URL string)

   gobotto.Parse(robotsContents string)

   gobotto.IsAllowed(userAgent string, route string, rules RobotsRules)


  This is how the whole proccess of checking if an agent is allowed to crawl a certain URL works:


    // First you've gotta fetch the robots.txt file contents
    // In order to do it you just gotta pass any URL to the Fetch function
    content, _ := gobotto.Fetch("https://example.com/path/to/any/page")

    // Now you need to parse the content you have fetched
    robots := gobotto.Parse(string(content))

    // Finally you just gotta check if an user-agent can crawl an URL given the rules we've just parsed
    allowed, _ := gobotto.isAllowed("agentName", "https://example.com/url/to/check", robots)
*/
package gobotto

import (
	// models module
	_ "github.com/lucasfcosta/gobotto/models"
	// robotsurl module
	_ "github.com/lucasfcosta/gobotto/robotsurl"
	// fetch module
	_ "github.com/lucasfcosta/gobotto/fetch"
	// parse module
	_ "github.com/lucasfcosta/gobotto/parse"
	// isallowed module
	_ "github.com/lucasfcosta/gobotto/isallowed"
)
