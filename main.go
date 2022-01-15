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
	var hasRunes bool
	var prevDash bool
	fn := func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			hasRunes = true
			prevDash = false
			return unicode.ToLower(r)
		}
		if hasRunes && !prevDash {
			prevDash = true
			return '-'
		}
		return -1
	}
	return strings.TrimRight(strings.Map(fn, text), "-")
}
