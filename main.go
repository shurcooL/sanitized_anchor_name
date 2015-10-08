// Package sanitized_anchor_name provides a func to create sanitized anchor names.
//
// Its logic can be reused by multiple packages to create interoperable anchor names
// and links to those anchors.
//
// At this time, it does not try to ensure that generated anchor names
// are unique, that responsibility falls on the caller.
package sanitized_anchor_name // import "github.com/shurcooL/sanitized_anchor_name"

import (
	"strings"
	"unicode"
)

// Create returns a sanitized anchor name for the given text.
func Create(text string) string {
	var anchorName []rune
	var lastWasDash = false

	for _, r := range []rune(strings.TrimSpace(text)) {
		switch {
		case r == ' ' || r == '-':
			if !lastWasDash {
				anchorName = append(anchorName, '-')
				lastWasDash = true
			}
		case unicode.IsLetter(r) || unicode.IsNumber(r):
			anchorName = append(anchorName, unicode.ToLower(r))
			lastWasDash = false
		default:
		}
	}

	// Special-case the situation where the entire string is spaces and
	// special chars.  At this point we'll have "-", which should be "".
	if len(anchorName) == 1 && anchorName[0] == '-' {
		return ""
	}

	return string(anchorName)
}
