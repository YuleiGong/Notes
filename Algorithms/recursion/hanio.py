#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-13 22:11:57
from __future__ import unicode_literals
from __future__ import absolute_import

"""
汉诺塔
ABC 三个圆盘。A圆盘上按照顶部小,底部小堆积
1:每次只能移动一个盘子
2:大盘子不能放到小盘子上
3:将A上的圆盘移动到B上
"""

def move_tower(n,A,B,C):
    """
    圆盘移动
    Args:
        n:圆盘数量
        A:第一根圆柱
        B:第二根圆柱
        C:第三根圆柱
    """
    if n == 1:
        print ("moving disk:{}->{}".format(A,C))
    else:
        move_tower(n-1,A,C,B) #移动到中间柱子
        print ("moving disk:{}->{}".format(A,C))
        move_tower(n-1,B,A,C) #中间注意依次移动到最后一根柱子



if __name__ == '__main__':
    move_tower(64,'A','B','C')
