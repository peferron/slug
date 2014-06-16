package slug

import (
	"testing"
)

var (
	tests = []struct {
		iS   string
		iL   int
		iSep string
		oS   string
	}{
		{"", 50, "-", ""},
		{" ", 50, "-", ""},
		{"a lovely day", 50, "-", "a-lovely-day"},
		{"a lovely day", 50, "123", "a123lovely123day"},
		{"A Lovely Day", 50, "-", "a-lovely-day"},
		{"a lovely day", 10, "-", "a-lovely-d"},
		{"a lovely day", 9, "-", "a-lovely"},
		{"a lovely day", 8, "-", "a-lovely"},
		{"wo ai 北京", 50, "-", "wo-ai-bei-jing"},
		{"wo ai 北京", 12, "-", "wo-ai-bei-ji"},
		{"wo ai 北京", 10, "-", "wo-ai-bei"},
		{"wo ai 北京", 10, "-", "wo-ai-bei"},
		{"wo ai 北京", 7, "-", "wo-ai-b"},
		{"wo ai 北京", 1, "-", "w"},
		{"Antonín Dvořák", 50, "-", "antonin-dvorak"},
		{"-a -- b__c", 50, "-", "a-b-c"},
		{"Schindler's List", 50, "-", "schindlers-list"},
		{"Schindler's List", 10, "-", "schindlers"},
	}
)

func TestSlug(t *testing.T) {
	for _, test := range tests {
		if out := Slug(test.iS, test.iL, test.iSep); out != test.oS {
			t.Errorf("For input string '%s' and maximum length %d and separator '%s', "+
				"expected output '%s' (len %d), was '%s' (len %d)",
				test.iS, test.iL, test.iSep,
				test.oS, len(test.oS), out, len(out))
		}
	}
}
