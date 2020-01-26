#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-01 16:47:39
from __future__ import unicode_literals
from __future__ import absolute_import

from stack import Stack

def par_checker(symbolstring):
    s = Stack()
    balanced = True
    index = 0

    while index < len(symbolstring) and balanced:
        symbol = symbolstring[index]
        if symbol in "([{":
            s.push(symbol)
        else:
            top = s.pop()
            if not matches(top, symbol):
                balanced = False

        index = index + 1

    if balanced and s.isEmpty():
        return True
    else:
        return False

def matches(open,close):
    opens = "([{"
    closers = ")]}"


    return opens.index(open) == closers.index(close)

if __name__ == '__main__':
    symbolstring = "({[)})"
    print (par_checker(symbolstring))


