// Package slug provides an utility for generating URL-friendly slugs.
//
// Example:
//
//	package main
//
//	import "github.com/peferron/slug"
//
//	func main() {
//		s := slug.Generate("Hello World!")
//		print(s) // Prints "hello-world"
//	}
package slug

import (
	"regexp"
	"strings"

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
	slug := strings.ToLower(decoded)

	// remove non-alphanumeric chars, except those in `chars`
	slug = replaceAll("[^a-z0-9" + regexp.QuoteMeta(chars) + "]", slug, "")

	// replace everything appearing in `chars` with seperator
	if len(chars) > 0 {
		slug = replaceAll("[" + regexp.QuoteMeta(chars) + "]+", slug, sep)
	}

	// remove prefixed seperator if present and enforce max length
	slug = strings.TrimPrefix(slug, sep)
	if len(slug) > l {
		slug = slug[:l]
	}

	return strings.TrimSuffix(slug, sep)
}

// replaceAll replaces everything in src matching regex with repl
func replaceAll(regex string, src string, repl string) string {
	re := regexp.MustCompile(regex)
	return re.ReplaceAllString(src, repl)
}
