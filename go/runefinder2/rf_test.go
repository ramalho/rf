package main

import (
	"os"
	"testing"
)

func Example_find() {
	find("cruzeiro")
	// Output:
	// U+20A2	₢	CRUZEIRO SIGN
	// (1 found)
}

func Example_find_two_results() {
	find("completion hexagram")
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

func Example_no_args() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"<executable-name>"}
	main()
	// Output:
	// Please give words to find or characters to get their names.
	// Example:
	// 	<executable-name> cat eyes
	// Use . or special characters in arguments to force character mode.
	// Example:
	// 	<executable-name> . 👋
}

func Example_listChars_liberation() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"<executable-name>", "40—A", "Liberação"}
	main()
	// Output:
	// U+0034	4	DIGIT FOUR
	// U+0030	0	DIGIT ZERO
	// U+2014	—	EM DASH
	// U+0041	A	LATIN CAPITAL LETTER A
	// U+0020	 	SPACE
	// U+004C	L	LATIN CAPITAL LETTER L
	// U+0069	i	LATIN SMALL LETTER I
	// U+0062	b	LATIN SMALL LETTER B
	// U+0065	e	LATIN SMALL LETTER E
	// U+0072	r	LATIN SMALL LETTER R
	// U+0061	a	LATIN SMALL LETTER A
	// U+00E7	ç	LATIN SMALL LETTER C WITH CEDILLA
	// U+00E3	ã	LATIN SMALL LETTER A WITH TILDE
	// U+006F	o	LATIN SMALL LETTER O
	// (14 found)
}

func Example_listChars_question_mark() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"<executable-name>", "What?"}
	main()
	// Output:
	// U+0057	W	LATIN CAPITAL LETTER W
	// U+0068	h	LATIN SMALL LETTER H
	// U+0061	a	LATIN SMALL LETTER A
	// U+0074	t	LATIN SMALL LETTER T
	// U+003F	?	QUESTION MARK
	// (5 found)
}

func Example_listChars_emoji() {
	listChars("Hi 👋")
	// Output:
	// U+0048	H	LATIN CAPITAL LETTER H
	// U+0069	i	LATIN SMALL LETTER I
	// U+0020	 	SPACE
	// U+1F44B	👋	WAVING HAND SIGN
	// (4 found)
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

func TestLikeUnicodeName(t *testing.T) {
	testCases := []struct {
		text string
		want bool
		name string
	}{
		{"hello-world name 42", true, "letters, hyphen, digits, space"},
		{"hello?", false, "question mark"},
		{"café", false, "non-ASCII character"},
		{"", false, "empty string"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := likeUnicodeName(tc.text)
			if got != tc.want {
				t.Errorf("likeUnicodeName(%q) = %v; want %v", tc.text, got, tc.want)
			}
		})
	}
}
