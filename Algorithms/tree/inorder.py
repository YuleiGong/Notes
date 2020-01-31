#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-30 20:46:32
from __future__ import unicode_literals
from __future__ import absolute_import

from ParseTree import buildParseTree

def inorder(tree):
    if tree != None:
        inorder(tree.getLeftChild())
        print (tree.getRootVal())
        inorder(tree.getRightChild())


if __name__ == '__main__':
    astr = "( 3 + ( 4 * 5 ) )"
    tree = buildParseTree(astr)
    inorder(tree)
