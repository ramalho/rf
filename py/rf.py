#!/usr/bin/env python3
"""
Test `find` with single result::

    >>> find("sign", "registered")  # doctest:+NORMALIZE_WHITESPACE
    U+00AE  ®   REGISTERED SIGN


Test `find` with two results::

    >>> find("dingbat", "zero")  # doctest:+NORMALIZE_WHITESPACE
    U+1F10B	🄋	DINGBAT CIRCLED SANS-SERIF DIGIT ZERO
    U+1F10C	🄌	DINGBAT NEGATIVE CIRCLED SANS-SERIF DIGIT ZERO

Test `main` with no words::

    >>> main([])
    Please provide one or more words to find.


Test `main` with two words::

    >>> main(['mark', 'check'])  # doctest:+NORMALIZE_WHITESPACE
    U+237B	⍻	NOT CHECK MARK
    U+2705	✅   WHITE HEAVY CHECK MARK
    U+2713	✓	CHECK MARK
    U+2714	✔	HEAVY CHECK MARK
    U+10102	𐄂	AEGEAN CHECK MARK
    U+1F5F8	🗸	LIGHT CHECK MARK


"""

import unicodedata
import sys

def find(*words):
    query = {w.upper() for w in words}
    for code in range(32, sys.maxunicode+1):
        char = chr(code)
        try:
            name = unicodedata.name(char)
        except ValueError:  # no such name
            continue
        if query <= set(name.split()):
            print(f'U+{ord(char):04X}\t{char}\t{name}')


def main(args):
	if len(args) < 1:
		print('Please provide one or more words to find.')
		return

	find(*args)

if __name__ == '__main__':
    main(sys.argv[1:])
