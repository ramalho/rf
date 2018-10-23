package main

import (
	"os"
)

func Example_find() {
	find("sign", "registered")
	// Output:
	// U+00AE	®	REGISTERED SIGN
	// (1 found)
}

func Example_find_two_results() {
	find("dingbat", "zero")
	// Output:
	// U+1F10B	🄋	DINGBAT CIRCLED SANS-SERIF DIGIT ZERO
	// U+1F10C	🄌	DINGBAT NEGATIVE CIRCLED SANS-SERIF DIGIT ZERO
	// (2 found)
}

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
	// (6 found)
}

func Example_no_args() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"<executable-name>"}
	main()
	// Output:
	// Please provide words to find.
}
