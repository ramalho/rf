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

    >>> find("chess", "queen", first=0x2654, last=0x2660)  # doctest:+NORMALIZE_WHITESPACE
    U+2655	♕	WHITE CHESS QUEEN
    U+265B	♛	BLACK CHESS QUEEN
    (2 found)

Test ``main`` with no words::

    >>> main([])
    Please provide words to find.
