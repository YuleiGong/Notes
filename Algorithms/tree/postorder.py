#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-30 20:34:36

from ParseTree import buildParseTree

def postorder(tree):
    if tree != 0:
        postorder(tree.getLeftChild())
        postorder(tree.getRightChild())
        print (tree.getRootVal())


if __name__ == '__main__':
    astr = "( 3 + ( 4 * 5 ) )"
    tree = buildParseTree(astr)
    tree.preorder()
