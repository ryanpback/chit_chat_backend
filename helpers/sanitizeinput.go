package helpers

import (
	"github.com/microcosm-cc/bluemonday"
)

// SanitizeInterface makes an interface with a string value a string to be passed into SanitizeInput
func SanitizeInterface(i interface{}) string {
	s := ConvertInterfaceToString(i)

	return SanitizeInput(s)
}

// SanitizeInput will sanitize all harmful html
// script, style, object, iframe, base, embed
func SanitizeInput(s string) string {
	p := bluemonday.UGCPolicy()

	return p.Sanitize(s)
}
