#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-02-21 13:03:22
from __future__ import unicode_literals
from __future__ import absolute_import

import multiprocessing
from multiprocessing import Barrier, Lock, Process
from time import time
from datetime import datetime

def test_with_barrier(synchronizer, serializer):
    name = multiprocessing.current_process().name
    synchronizer.wait()
    now = time()

    with serializer:
        print ("process {} -----> {}".format(name,datetime.fromtimestamp(now)))



if __name__ == '__main__':
    synchronizer = Barrier(3) #管理的进程数
    serializer = Lock()
    Process(name="p1 - test_with_barrier", target=test_with_barrier,args=(synchronizer,serializer)).start()
    Process(name="p2 - test_with_barrier", target=test_with_barrier,args=(synchronizer,serializer)).start()
    Process(name="p3 - test_with_barrier", target=test_with_barrier,args=(synchronizer,serializer)).start()





