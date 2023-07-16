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
	os.Args = []string{"<executable-name>", "great", "hexagram"}
	main()
	// Output:
	// U+4DCD	䷍	HEXAGRAM FOR GREAT POSSESSION
	// U+4DD9	䷙	HEXAGRAM FOR GREAT TAMING
	// U+4DDB	䷛	HEXAGRAM FOR GREAT PREPONDERANCE
	// U+4DE1	䷡	HEXAGRAM FOR GREAT POWER
	// (4 found)
}

func Example_no_args() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"<executable-name>"}
	main()
	// Output:
	// Please provide words to find.
}
