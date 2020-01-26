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
        if symbol == "(":
            s.push(symbol)
        else:
            if s.isEmpty():
                balanced = False
            else:
                s.pop()
        index = index + 1

    if balanced and s.isEmpty():
        return True
    else:
        return False

if __name__ == '__main__':
    symbolstring = list(("((()))"))
    print (par_checker(symbolstring))


