#!/usr/bin/env python3
import sys
import unicodedata


def find(*words):
    query = {w.upper() for w in words}
    for code in range(32, sys.maxunicode + 1):
        char = chr(code)
        name = unicodedata.name(char, "")
        if not name:
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
