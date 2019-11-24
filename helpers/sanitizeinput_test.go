package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeInputScriptTagRemoveTag(t *testing.T) {
	assert := assert.New(t)
	htmlString := "<script>alert('hello world')</script>"
	sanitizedString := SanitizeInput(htmlString)

	assert.Equal("", sanitizedString)
}

func TestSanitizeInputScriptTagLeaveTextAroundIt(t *testing.T) {
	assert := assert.New(t)
	htmlString := "HELLO <script>alert('hello world')</script>WORLD"
	sanitizedString := SanitizeInput(htmlString)

	assert.Equal("HELLO WORLD", sanitizedString)
}

func TestSanitizeInputLeaveATag(t *testing.T) {
	assert := assert.New(t)
	htmlString := "<a href='example.com'>Click Here</a>"
	sanitizedString := SanitizeInput(htmlString)

	// Doesn't actually escape quotes, only here because it's a double quote within a double quote
	assert.Equal("<a href=\"example.com\" rel=\"nofollow\">Click Here</a>", sanitizedString)
}

func TestSanitizeInputRemoveATagIfHrefIsJS(t *testing.T) {
	assert := assert.New(t)
	htmlString := "<a href=\"javascript:alert('XSS1')\" onmouseover=\"alert('XSS2')\">XSS<a>"
	sanitizedString := SanitizeInput(htmlString)

	assert.Equal("XSS", sanitizedString)
}

func TestSanitizeInterface(t *testing.T) {
	assert := assert.New(t)
	var payload = map[string]interface{}{
		"html": "<a href=\"javascript:alert('XSS1')\" onmouseover=\"alert('XSS2')\">XSS<a>",
	}

	sanitizedString := SanitizeInterface(payload["html"])

	assert.Equal("XSS", sanitizedString)
}
