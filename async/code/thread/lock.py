#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-02-18 14:35:37
from __future__ import unicode_literals
from __future__ import absolute_import

import threading

share_resource_with_lock = 0
share_resource_with_no_lock = 0
COUNT = 100000
share_resource_lock = threading.Lock()

def increment_with_lock():
    global share_resource_with_lock
    for i in range(COUNT):
        share_resource_lock.acquire()
        share_resource_with_lock += 1
        share_resource_lock.release()

def decrement_with_lock():
    global share_resource_with_lock
    for i in range(COUNT):
        share_resource_lock.acquire()
        share_resource_with_lock -= 1
        share_resource_lock.release()

def increment_without_lock():
    global share_resource_with_no_lock
    for i in range(COUNT):
        share_resource_with_no_lock += 1

def decrement_without_lock():
    global share_resource_with_no_lock
    for i in range(COUNT):
        share_resource_with_no_lock -= 1

if __name__ == "__main__":
    t1 = threading.Thread(target=increment_with_lock)
    t2 = threading.Thread(target=decrement_with_lock)
    t3 = threading.Thread(target=increment_without_lock)
    t4 = threading.Thread(target=decrement_without_lock)

    t1.start()
    t2.start()
    t3.start()
    t4.start()

    t1.join()
    t2.join()
    t3.join()
    t4.join()

    print ("the value with no lock is {}".format(share_resource_with_no_lock))
    print ("the value with lock is {}".format(share_resource_with_lock))


