// Package slug provides an utility for generating URL-friendly slugs.
//
// Example:
//
//	package main
//
//	import "github.com/peferron/go-slug"
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

// Slug returns a slug generated from s, with the maximum length l and the separator sep.
func Slug(s string, l int, sep string) string {
	decoded := unidecode.Unidecode(s)
	slug := ""

	for _, d := range decoded {
		r := convert(d, sep)

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
func convert(r rune, sep string) string {
	switch {
	case 'a' <= r && r <= 'z' || '0' <= r && r <= '9':
		return string(r)
	case 'A' <= r && r <= 'Z':
		return string(unicode.ToLower(r))
	case r == ' ' || r == ',' || r == '.' || r == '/' || r == '\\' || r == '-' || r == '_' ||
		r == '=':
		return sep
	}
	return ""
}
