#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-02-18 13:56:19
from __future__ import unicode_literals
from __future__ import absolute_import

import threading
import time

def first_function():
    print (threading.currentThread().getName() + " is Strating")
    time.sleep(2)
    print (threading.currentThread().getName() + " is Exiting")
    return

def second_function():
    print (threading.currentThread().getName() + " is Strating")
    time.sleep(2)
    print (threading.currentThread().getName() + " is Exiting")
    return


def third_function():
    print (threading.currentThread().getName() + " is Strating")
    time.sleep(2)
    print (threading.currentThread().getName() + " is Exiting")
    return

if __name__ == '__main__':
    t1 = threading.Thread(name='first_function',target=first_function)
    t2 = threading.Thread(name='second_function',target=second_function)
    t3 = threading.Thread(target=third_function)
    t1.start()
    t2.start()
    t3.start()

    t1.join()
    t2.join()
    t3.join()

