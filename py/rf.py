#!/usr/bin/env python3
import sys
import unicodedata

FIRST, LAST = ord(" "), sys.maxunicode


def find(*words):
    query = {w.upper() for w in words}
    count = 0
    for code in range(FIRST, LAST + 1):
        char = chr(code)
        name = unicodedata.name(char, "")
        if not name:
            continue
        if query <= set(name.split()):
            print(f"U+{code:04X}\t{char}\t{name}")
            count += 1
    print(f"({count} found)")


def main(args):
    if len(args) < 1:
        print("Please provide words to find.")
        return

    find(*args)


if __name__ == "__main__":
    main(sys.argv[1:])
