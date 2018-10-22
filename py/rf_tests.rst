Doctests for ``rf.py``
======================

How to run the tests
----------------------

Run the ``doctest`` module from the command line::

    $ python3 -m doctest rf_tests.rst


Tests
-----

Import functions for testing::

    >>> from rf import find, main

Test ``find`` with single result::

    >>> find("sign", "registered")  # doctest:+NORMALIZE_WHITESPACE
    U+00AE  ®   REGISTERED SIGN


Test ``find`` with two results::

    >>> find("dingbat", "zero")  # doctest:+NORMALIZE_WHITESPACE
    U+1F10B	🄋	DINGBAT CIRCLED SANS-SERIF DIGIT ZERO
    U+1F10C	🄌	DINGBAT NEGATIVE CIRCLED SANS-SERIF DIGIT ZERO

Test ``main`` with no words::

    >>> main([])
    Please provide one or more words to find.


Test ``main`` with two words::

    >>> main(['mark', 'check'])  # doctest:+NORMALIZE_WHITESPACE
    U+237B	⍻	NOT CHECK MARK
    U+2705	✅   WHITE HEAVY CHECK MARK
    U+2713	✓	CHECK MARK
    U+2714	✔	HEAVY CHECK MARK
    U+10102	𐄂	AEGEAN CHECK MARK
    U+1F5F8	🗸	LIGHT CHECK MARK
