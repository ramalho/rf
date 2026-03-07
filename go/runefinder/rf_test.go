package main

import (
	"os"
	"slices"
	"testing"
	"unicode"
)

const first, last = ' ', unicode.MaxRune

func TestTokenize(t *testing.T) {
	testCases := []struct {
		input string
		want  []string
		name  string
	}{
		{"hello", []string{"hello"}, "single word"},
		{"hello new world", []string{"hello", "new", "world"}, "3 words"},
		{"hello-world", []string{"hello", "world"}, "hyphenated word"},
		{"multi-hyphen-word", []string{"multi", "hyphen", "word"}, "multiple hyphens"},
		{"word1 word2-word3", []string{"word1", "word2", "word3"}, "mixed spaces and hyphens"},
		{"", []string{}, "empty string"},
		{"   ", []string{}, "whitespace only"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tokenize(tc.input)
			if !slices.Equal(got, tc.want) {
				t.Errorf("got %v; want %v", got, tc.want)
			}
		})
	}
}

func TestContainsAll(t *testing.T) {
	testCases := []struct {
		list  []string
		items []string
		want  bool
		name  string
	}{
		{[]string{}, []string{"B"}, false, "1 item, empty list"},
		{[]string{"A"}, []string{"B"}, false, "1 item, not found"},
		{[]string{"B"}, []string{"B"}, true, "1 item, found"},
		{[]string{"A", "B"}, []string{"B"}, true, "1 item, found last"},
		{[]string{}, []string{"B", "A"}, false, "2 items, empty list"},
		{[]string{"A"}, []string{"B", "A"}, false, "2 items, B not found"},
		{[]string{"A"}, []string{"B", "A"}, false, "2 items, A not found"},
		{[]string{"A", "B"}, []string{"B", "A"}, true, "2 items, found"},
		{[]string{"A", "B", "C"}, []string{"B", "A"}, true, "2 items, found with extra"},
		{[]string{}, []string{}, true, "0 item, empty list"},
		{[]string{"A"}, []string{}, true, "0 item, any list"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := containsAll(tc.list, tc.items)
			if got != tc.want {
				t.Errorf("got %v; want %v", got, tc.want)
			}
		})
	}
}

func Example() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"<executable-name>", "apl", "tilde"}
	main()
	// Output:
	// U+2368	⍨	APL FUNCTIONAL SYMBOL TILDE DIAERESIS
	// U+236B	⍫	APL FUNCTIONAL SYMBOL DEL TILDE
	// U+236D	⍭	APL FUNCTIONAL SYMBOL STILE TILDE
	// U+2371	⍱	APL FUNCTIONAL SYMBOL DOWN CARET TILDE
	// U+2372	⍲	APL FUNCTIONAL SYMBOL UP CARET TILDE
	// (5 found)
}

func Example_find() {
	find("cruzeiro", first, last)
	// Output:
	// U+20A2	₢	CRUZEIRO SIGN
	// (1 found)
}

func Example_find_two_results() {
	find("completion hexagram", first, last)
	// Output:
	// U+4DFE	䷾	HEXAGRAM FOR AFTER COMPLETION
	// U+4DFF	䷿	HEXAGRAM FOR BEFORE COMPLETION
	// (2 found)
}

func Example_find_hyphenated_name() {
	find("hyphen", ' ', '\x80')
	// Output:
	// U+002D	-	HYPHEN-MINUS
	// (1 found)
}

func Example_no_args() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"<executable-name>"}
	main()
	// Output:
	// Please provide words to find.
	// Example:
	// <executable-name> cat eyes
}
