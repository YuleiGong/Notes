#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-30 14:27:19

"""
构建标准二叉树
"""

def BinaryTree(r):
    """
    添加二叉树的新节点
    """
    return [r, [], []]

def insertLeft(root, newBranch):
    """
    插入左子树
    """
    #如果左子树有内容,需要将旧的左子树作为新节点的左子树
    t = root.pop(1)
    if len(t) > 1:
        root.insert(1, [newBranch,t,[]])
    else:
        root.insert(1,[newBranch,[],[]])

    return root

def insertRight(root, newBranch):
    """
    插入右子树
    """
    t = root.pop(2)
    if len(t) > 1:
        root.insert(2,[newBranch,[],t])
    else:
        root.insert(2, [newBranch,[],[]])

    return root

def getRootVal(root):
    return root[0]

def setRootVal(root, newVal):
    root[0] = newVal

def getLeftChild(root):
    return root[1]

def getRightChild(root):
    return root[2]




if __name__ == '__main__':
    pass

