#!/usr/bin/env python3
import sys
import unicodedata

FIRST, LAST = ord(' '), sys.maxunicode


def tokenize(text):
    return text.replace('-', ' ').split()


def find(text, first=FIRST, last=LAST):
    query = set(tokenize(text.upper()))
    count = 0
    for code in range(first, last + 1):
        char = chr(code)
        name = unicodedata.name(char, '')
        if name and query.issubset(tokenize(name)):
            print(f'U+{code:04X}\t{char}\t{name}')
            count += 1
    print(f'({count} found)')


def main(args):
    if args:
        find(' '.join(args))
    else:
        print('Please provide words to find.')


if __name__ == '__main__':
    main(sys.argv[1:])
