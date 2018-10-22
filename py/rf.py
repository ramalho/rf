#!/usr/bin/env python3
import unicodedata
import sys


def find(*words):
    query = {w.upper() for w in words}
    for code in range(32, sys.maxunicode + 1):
        char = chr(code)
        try:
            name = unicodedata.name(char)
        except ValueError:  # no such name
            continue
        if query <= set(name.split()):
            print(f"U+{ord(char):04X}\t{char}\t{name}")


def main(args):
    if len(args) < 1:
        print("Please provide one or more words to find.")
        return

    find(*args)


if __name__ == "__main__":
    main(sys.argv[1:])
