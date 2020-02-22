#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-02-21 11:12:52
from __future__ import unicode_literals
from __future__ import absolute_import

import multiprocessing

def foo(i):
    print ("called function in process: {}".format(i))
    return


if __name__ == '__main__':
    jobs = []
    for i in range(5):
        p = multiprocessing.Process(target=foo, args=(i,))
        jobs.append(p)
        p.start()
        p.join()

