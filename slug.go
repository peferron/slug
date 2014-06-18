// Package slug provides an utility for generating URL-friendly slugs.
//
// Example:
//
//	package main
//
//	import "github.com/peferron/slug"
//
//	func main() {
//		s := slug.Slug("Hello World!", 50, "-")
//		print(s) // Prints "hello-world"
//	}
package slug

import (
	"strings"
	"unicode"

	"github.com/fiam/gounidecode/unidecode"
)

// Generate returns a slug generated from s, with sensible default parameters.
func Generate(s string) string {
	return GenerateWith(s, ` ,./\-_=~|()[]{}:;?!<>`, "-", 100)
}

// GenerateWith returns a slug generated from s, with all Unicode code points in chars replaced by
// the separator sep, and the final maximum length l.
func GenerateWith(s, chars, sep string, l int) string {
	decoded := unidecode.Unidecode(s)
	slug := ""

	for _, d := range decoded {
		r := convert(d, chars, sep)

		switch r {
		case "":
			continue
		case sep:
			if len(slug) > 0 && !strings.HasSuffix(slug, sep) {
				slug += sep
			}
		default:
			slug += string(r)
		}

		if len(slug) >= l {
			break
		}
	}

	return strings.TrimSuffix(slug, sep)
}

// convert returns the slug representation of r, using the separator sep.
func convert(r rune, chars, sep string) string {
	switch {
	case 'a' <= r && r <= 'z' || '0' <= r && r <= '9':
		return string(r)
	case 'A' <= r && r <= 'Z':
		return string(unicode.ToLower(r))
	case strings.ContainsRune(chars, r):
		return sep
	}
	return ""
}
