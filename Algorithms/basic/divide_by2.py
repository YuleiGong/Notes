#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-01 20:18:21
from __future__ import unicode_literals
from __future__ import absolute_import

from stack import Stack

def divide_by2(dec_number):
    remstack = Stack()

    while dec_number > 0:
        rem = dec_number % 2 #余数部分
        remstack.push(rem)
        dec_number = dec_number // 2 #整数部分

    bin_string = ""
    while not remstack.isEmpty():
        bin_string = bin_string + str(remstack.pop())

    return bin_string

def divide_converter(dec_number, base):
    """
    十进制数转换成任意进制数字
    Args:
        dec_number:十进制数
        base:转换基数
    Returns:
        返回base进制的数
    """
    digits = "0123456789ABCDEF"

    remstack = Stack()

    while dec_number > 0:
        rem = dec_number % base #余数部分
        remstack.push(rem)
        dec_number = dec_number // base #整数部分

    new_string = ""
    while not remstack.isEmpty():
        new_string = new_string + digits[remstack.pop()]

    return new_string





if __name__ == '__main__':
    #print (divide_by2(233))
    print (divide_converter(10,2))
    #print (divide_converter.__doc__)

