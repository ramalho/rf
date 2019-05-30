# Rune finders

`rf`: utility for locating Unicode characters by name. Simplest sensible implementations in different languages.

## What it does

Given one or more words, `rf` should return a list of all Unicode characters that have those words in their names. Only complete words should be matched, and the match is case-insenstive. For example:

```
$ ./rf cat eyes
U+1F638	😸	GRINNING CAT FACE WITH SMILING EYES
U+1F63B	😻	SMILING CAT FACE WITH HEART-SHAPED EYES
U+1F63D	😽	KISSING CAT FACE WITH CLOSED EYES
(3 found)
```

## Where to get the data

The standard names of Unicode characters can be found in the [UnicodeData.txt](http://www.unicode.org/Public/UNIDATA/UnicodeData.txt) text file (32K lines, 1.7 Mb).

Some programming languages provide standard packages with that information. For example, in Python, there is the [unicodedata](https://docs.python.org/3/library/unicodedata.html) package, and Go has the semi-official [runenames](https://godoc.org/golang.org/x/text/unicode/runenames) package.

Using the official [UnicodeData.txt](http://www.unicode.org/Public/UNIDATA/UnicodeData.txt) has two advantages. First, you get it directly from the Unicode Consortium, so it is usually more current than the data packaged in a library, including the newest characters and emojis. Second, it provides alternative names for some characters, which sometimes are better than the standard name. For example, the standard name for `\` is "reverse solidus", but [UnicodeData.txt](http://www.unicode.org/Public/UNIDATA/UnicodeData.txt) also offers "backslash" as an alternative.
