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
