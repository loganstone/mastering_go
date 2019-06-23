package ipv4

import (
	"regexp"
)

const part = "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
const grammar = part + "\\." + part + "\\." + part + "\\." + part

var r = regexp.MustCompile(grammar)

// Find .
func Find(s string) string {
	return r.FindString(s)
}
