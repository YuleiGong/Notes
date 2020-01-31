#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-30 20:17:05

from ParseTree import buildParseTree


def preorder(tree):
    """
    前序遍历
        Args:
            tree:二叉树对象
    """
    if tree:
        print (tree.getRootVal())
        preorder(tree.getLeftChild())
        preorder(tree.getRightChild())

if __name__ == '__main__':
    astr = "( 3 + ( 4 * 5 ) )"
    tree = buildParseTree(astr)
    tree.preorder()
    #preorder(tree)

