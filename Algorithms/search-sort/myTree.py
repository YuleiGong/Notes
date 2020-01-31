#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-30 14:15:46

myTree = [
    'a',#根节点
    ['b',['d',[],[]],['e',[],[]]],#左字树
    ['c',['f',[],[]]],#右子树
]

if __name__ == '__main__':
    print ("根节点:{}".format(myTree[0])) #根节点
    print ("左子树:{}".format(myTree[1])) #左子树
    print ("右子树:{}".format(myTree[2])) #右子树
