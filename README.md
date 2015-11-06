Package `scriptsources` is a simple utility that extracts resources referenced inside of HTML script tags. Given a URL, it'll retrieve the page and comb through the body of the response finding all resources specified in the  `<script>` tags' `src` attributes.

[![GoDoc](https://godoc.org/github.com/jboursiquot/scriptsources?status.svg)](http://godoc.org/github.com/jboursiquot/scriptsources)

### Usage

```go
package main

import (
	"github.com/jboursiquot/scriptsources"

	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://jboursiquot.com")
	checkErr(err)
	defer resp.Body.Close()

	resources := scriptsources.ExtractResources(resp.Body)

	fmt.Println(resources)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
```

Your output should be something like this:

[http://jboursiquot.com/wp-includes/js/jquery/jquery.js?ver=1.11.3 http://jboursiquot.com/wp-includes/js/jquery/jquery-migrate.min.js?ver=1.2.1 http://jboursiquot.com/wp-content/themes/zeebizzcard/includes/js/jquery.cycle.all.min.js?ver=4.3.1 http://jboursiquot.com/wp-content/plugins/google-analyticator/external-tracking.min.js?ver=6.4.9 //platform.twitter.com/oct.js]

Enjoy.
