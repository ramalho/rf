# Rune finder script in Elixir

The `rf.exs` script allows finding Unicode characters by name:

```
$ ./rf.exs face cat
U+1F431	🐱	CAT FACE
U+1F638	😸	GRINNING CAT FACE WITH SMILING EYES
U+1F639	😹	CAT FACE WITH TEARS OF JOY
U+1F63A	😺	SMILING CAT FACE WITH OPEN MOUTH
U+1F63B	😻	SMILING CAT FACE WITH HEART-SHAPED EYES
U+1F63C	😼	CAT FACE WITH WRY SMILE
U+1F63D	😽	KISSING CAT FACE WITH CLOSED EYES
U+1F63E	😾	POUTING CAT FACE
U+1F63F	😿	CRYING CAT FACE
U+1F640	🙀	WEARY CAT FACE
```

To operate, `rf.exs` reads the `UnicodeData.txt` file in this directory.

To get the most up-to-date version of `UnicodeData.txt`,
download it from the official source:

```
$ curl -O ftp://ftp.unicode.org/Public/UNIDATA/UnicodeData.txt
```
