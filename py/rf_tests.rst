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
    (1 found)


Test ``find`` with two results::

    >>> find("chess", "queen")  # doctest:+NORMALIZE_WHITESPACE
    U+2655	♕	WHITE CHESS QUEEN
    U+265B	♛	BLACK CHESS QUEEN
    (2 found)

Test ``main`` with no words::

    >>> main([])
    Please provide words to find.


Test ``main`` with two words::

    >>> main(['chess', 'black'])  # doctest:+NORMALIZE_WHITESPACE
    U+265A	♚	BLACK CHESS KING
    U+265B	♛	BLACK CHESS QUEEN
    U+265C	♜	BLACK CHESS ROOK
    U+265D	♝	BLACK CHESS BISHOP
    U+265E	♞	BLACK CHESS KNIGHT
    U+265F	♟	BLACK CHESS PAWN
    (6 found)
