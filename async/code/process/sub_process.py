#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-02-21 11:50:11
from __future__ import unicode_literals
from __future__ import absolute_import

import multiprocessing

class MyProcess(multiprocessing.Process):

    def __init__(self):
        super(MyProcess, self).__init__()

    def run(self):
        print ("called run method in process: {}".format(self.name))


if __name__ == '__main__':
    jobs = []
    for i in range(5):
        p = MyProcess()
        jobs.append(p)
        p.start()
    [p.join() for p in jobs]

