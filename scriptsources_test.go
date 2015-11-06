package scriptsources

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestExtractResources(t *testing.T) {
	reader := strings.NewReader(`<html><head><script src="//ajax.googleapis.com/ajax/libs/jquery/1.11.0/jquery.min.js" type="text/javascript"></script></head><body>
    <a href="somereference.html">Some Reference</a>
    <script src="https://ajax.googleapis.com/ajax/libs/angularjs/1.3.15/angular.min.js"></script>
    <script src="//somedomain.com/somelib.js"></script>
    </body>
    </html>`)

	resources := ExtractResources(reader)

	assert.Equal(t, len(resources), 3, "there should be 3 resources found")

	assert.Equal(t, resources[0], "//ajax.googleapis.com/ajax/libs/jquery/1.11.0/jquery.min.js", "expected a specific resource at position 0")

	assert.Equal(t, resources[2], "//somedomain.com/somelib.js", "expected a specific resource at the last position")
}
