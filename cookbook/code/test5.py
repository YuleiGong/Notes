#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2019-07-18 11:34:25
from __future__ import unicode_literals
from __future__ import absolute_import


from queue import Queue
from threading import Thread

_sentinel = object()

def produce(out_q):
    while runing:
        out_q.put(data)
    out_q.put(data)


def consumer(in_q):
    while True:
        data = in_q.get()
        if data is _sentinel:
            in_q.put(_sentinel)
            break


if __name__ == '__main__':
    q = Queue()
    t1 = Thread(target=consumer, args=(q,))
    t2 = Thread(target=produce, args=(q,))
    t1.start()
    t2.start()
