package slug

import "testing"

func TestGenerate(t *testing.T) {
	tests := []struct {
		iS string
		oS string
	}{
		{"", ""},
		{" ", ""},
		{"a lovely day", "a-lovely-day"},
		{"A Lovely Day", "a-lovely-day"},
		{"---a lovely day---", "a-lovely-day"},
		{"wo ai 北京", "wo-ai-bei-jing"},
		{"Antonín Dvořák", "antonin-dvorak"},
		{"Schindler's List", "schindlers-list"},
	}

	for _, test := range tests {
		if out := Generate(test.iS); out != test.oS {
			t.Errorf("For input string %q, expected output %q (len %d), was %q (len %d)",
				test.iS, test.oS, len(test.oS), out, len(out))
		}
	}
}

func TestGenerateWith(t *testing.T) {
	tests := []struct {
		iS     string
		iChars string
		iSep   string
		iL     int
		oS     string
	}{
		{"a lovely day", "", "-", 100, "alovelyday"},
		{"a lovely day", " ", "*123*", 100, "a*123*lovely*123*day"},
		{"a lovely day", " ", "_", 10, "a_lovely_d"},
		{"a lovely day", " ", "_", 9, "a_lovely"},
		{"a lovely day", " ", "_", 8, "a_lovely"},
		{"wo ai 北京", " ", "-", 12, "wo-ai-bei-ji"},
		{"wo ai 北京", " ", "-", 10, "wo-ai-bei"},
		{"wo ai 北京", " ", "-", 7, "wo-ai-b"},
		{"wo ai 北京", " ", "-", 1, "w"},
		{"-a -- b__c", " -", "-", 50, "a-bc"},
		{"Schindler's List", " ", "-", 100, "schindlers-list"},
		{"Schindler's List", " ", "-", 10, "schindlers"},
		{"Schindler's List", " '", "-", 100, "schindler-s-list"},
		{"Schindler's List", " '", "-", 10, "schindler"},
	}

	for _, test := range tests {
		if out := GenerateWith(test.iS, test.iChars, test.iSep, test.iL); out != test.oS {
			t.Errorf("For input string %q, chars %q, separator %q and maximum length %d, "+
				"expected output %q (len %d), was %q (len %d)",
				test.iS, test.iChars, test.iSep, test.iL,
				test.oS, len(test.oS), out, len(out))
		}
	}
}
