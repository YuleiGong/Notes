#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-02-22 11:30:44
from __future__ import unicode_literals
from __future__ import absolute_import

import multiprocessing

def worker(dictionary, key, item):
    dictionary[key] = item
    print ("key={} value={}".format(key, item))


if __name__ == '__main__':
    mgr = multiprocessing.Manager()
    dictionary = mgr.dict()
    jobs = [multiprocessing.Process(target=worker, args=(dictionary,i,i*2)) for i in range(10)]
    for j in jobs

