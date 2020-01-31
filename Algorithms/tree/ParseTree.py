#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-30 16:57:11

"""
构造解析树
"""
from stack import Stack
from binary_tree1 import BinaryTree
import operator



def buildParseTree(fpexp):
    """
    * 如果当前标记是(，就为当前节点添加一个左子节点，并下沉至该子节点;
    * 如果当前标记在列表['+', '-', '/', '\*']中,就将当前节点的值设为当前标记对应的运算符;为当前节点添加一个右子节点，并下沉至该子节点;
    * 如果当前标记是数字，就将当前节点的值设为这个数并返回至父节点;
    * 如果当前标记是)，就跳到当前节点的父节点。
    Args:
        fpexp:表达式,注意使用空格分隔每个字符 exp: "( 3 + ( 4 * 5 ) )"
    Returns:
        eTree:二叉树对象
    """
    fplist = fpexp.split()
    pStack = Stack()
    eTree = BinaryTree('')
    pStack.push(eTree)
    currentTree = eTree

    for i in fplist:
        if i == '(':
            currentTree.insertLeft('')
            pStack.push(currentTree)
            currentTree = currentTree.getLeftChild()
        elif i not in '+-*/)': #数字
            currentTree.setRootVal(eval(i))
            parent = pStack.pop()
            currentTree = parent
        elif i in '+-*/': 
            currentTree.setRootVal(i)
            currentTree.insertRight('')
            pStack.push(currentTree)
            currentTree = currentTree.getRightChild()
        elif i == ')':
            currentTree = pStack.pop()
        else:
            raise ValueError("Unknow Operator: " + i)

    return eTree

def evalute(parseTree):
    """
    计算二叉解析树
    Args:
        parseTree:表达式二叉树
    Returns:
        计算结果
    """
    opers = {'+':operator.add,'-':operator.sub,
             '*':operator.mul,'/':operator.truediv
    }
    leftC = parseTree.getLeftChild()
    rightC = parseTree.getRightChild()

    if leftC and rightC:
        fn = opers[parseTree.getRootVal()]
        return fn(evalute(leftC), evalute(rightC))
    else:
        return parseTree.getRootVal()




if __name__ == '__main__':
    astr = "( 3 + ( 4 * 5 ) )"
    tree = buildParseTree(astr)
    print (evalute(tree))

