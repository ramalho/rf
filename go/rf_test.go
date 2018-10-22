package main

import (
	"os"
)

func Example() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"<executable-name>", "mark", "check"}
	main()
	// Output:
	// U+237B	⍻	NOT CHECK MARK
	// U+2705	✅	WHITE HEAVY CHECK MARK
	// U+2713	✓	CHECK MARK
	// U+2714	✔	HEAVY CHECK MARK
	// U+10102	𐄂	AEGEAN CHECK MARK
	// U+1F5F8	🗸	LIGHT CHECK MARK
}

func Example_no_args() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"<executable-name>"}
	main()
	// Output:
	// Please provide one or more words to find.
}
func ExampleFindOne() {
	find("sign", "registered")
	// Output:
	// U+00AE	®	REGISTERED SIGN
}

func ExampleFindTwo() {
	find("dingbat", "zero")
	// Output:
	// U+1F10B	🄋	DINGBAT CIRCLED SANS-SERIF DIGIT ZERO
	// U+1F10C	🄌	DINGBAT NEGATIVE CIRCLED SANS-SERIF DIGIT ZERO
}

func FutureExampleFindHyphenated() {
	find("plus", "minus")
	// Output:
	// U+00B1	±	PLUS-MINUS SIGN (PLUS-OR-MINUS SIGN)
	// U+2213	∓	MINUS-OR-PLUS SIGN
}
