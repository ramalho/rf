# Copilot Instructions for RF (Rune Finder)

## Project Overview
This is a polyglot educational project implementing `rf` - a Unicode character search utility - in multiple languages (Python, Go, Elixir, Kotlin, JavaScript). Each implementation searches Unicode characters by name keywords and displays matching results with code points.

## Architecture & Data Flow
- **Core Logic**: Parse Unicode character names, tokenize search queries, match complete words (case-insensitive)
- **Data Sources**: `UnicodeData.txt` files contain semicolon-separated Unicode character data
- **Output Format**: `U+XXXX\t{char}\t{NAME}` followed by count summary
- **Key Algorithm**: Word tokenization replaces hyphens with spaces, then splits on whitespace

## Language-Specific Patterns

### Python (`py/rf.py`)
- Uses `unicodedata` standard library for character names
- Doctests in separate `.rst` file (`rf_tests.rst`)
- Run tests: `python3 -m doctest rf_tests.rst` or `./test.sh`

### Go (`go/runefinder/` and `go/runefinder2/`)
- Uses `golang.org/x/text/unicode/runenames` package
- Two versions: basic (`runefinder`) and enhanced (`runefinder2` with regex validation)
- Go example tests in `rf_test.go`
- Run tests: `go test`

### Elixir (`elixir/rf.exs`)
- Executable script reads local `UnicodeData.txt` file
- Uses `MapSet` for efficient word matching
- Pipe-based functional style with pattern matching

### Kotlin (`kotlin/rf.kt`)
- Uses Java's `Character.getName()` method
- Follows Go's API pattern with varargs for range specification

## Development Workflows

### Testing Patterns
- **Python**: Doctest format with `+NORMALIZE_WHITESPACE` directive
- **Go**: Example tests and standard `testing` package
- **Elixir**: Run directly as executable script
- **Data**: Reduced test datasets (e.g., `TestData.txt`) for focused testing

### Data Management
- Download fresh Unicode data: `curl -O ftp://ftp.unicode.org/Public/UNIDATA/UnicodeData.txt`
- Language-specific subsets: ASCII (`AsciiData.txt`), Latin-1 (`Latin1Data.txt`)
- Format: `{code};{name};{category};...` (semicolon-separated, name is field 2)

### Key Implementation Details
- **Tokenization**: Always replace hyphens with spaces before splitting
- **Matching**: Use set operations or contains-all logic for word matching
- **Range Support**: Optional first/last parameters for character code ranges
- **Error Handling**: Skip characters without names, provide usage messages

## Project Conventions
- Each language lives in its own directory with minimal dependencies
- Consistent CLI interface: `./rf word1 word2 ...`
- Shared test cases: "cat eyes", "chess queen", "sign registered"
- Unicode code points formatted as uppercase hex with U+ prefix