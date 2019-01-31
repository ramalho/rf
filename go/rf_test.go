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
	find("chess", "queen")
	// Output:
	// U+2655	♕	WHITE CHESS QUEEN
	// U+265B	♛	BLACK CHESS QUEEN
	// (2 found)
}

func Example() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"<executable-name>", "chess", "black"}
	main()
	// Output:
	// U+265A	♚	BLACK CHESS KING
	// U+265B	♛	BLACK CHESS QUEEN
	// U+265C	♜	BLACK CHESS ROOK
	// U+265D	♝	BLACK CHESS BISHOP
	// U+265E	♞	BLACK CHESS KNIGHT
	// U+265F	♟	BLACK CHESS PAWN
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
