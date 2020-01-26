#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-01-14 21:41:51
from __future__ import unicode_literals
from __future__ import absolute_import

from queue import Queue

def hot_potato(namelist, num):
    simqueue = Queue()
    for name in namelist:
        simqueue.enqueue(name)

    while simqueue.size() > 1:
        for i in range(num):
            simqueue.enqueue(simqueue.dequeue())
        #num次循环完成后,出局的数据
        simqueue.dequeue()

    return simqueue.dequeue()

if __name__ == '__main__':
    print (hot_potato(["Bill", "David", "Susan", "Jane", "Kent", "Brad"],7))

