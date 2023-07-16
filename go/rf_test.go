package main

import (
	"os"
)

func Example_find() {
	find("cruzeiro")
	// Output:
	// U+20A2	₢	CRUZEIRO SIGN
	// (1 found)
}

func Example_find_two_results() {
	find("hexagram", "completion")
	// Output:
	// U+4DFE	䷾	HEXAGRAM FOR AFTER COMPLETION
	// U+4DFF	䷿	HEXAGRAM FOR BEFORE COMPLETION
	// (2 found)
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
	// Please provide words to find.
}
