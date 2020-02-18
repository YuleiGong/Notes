#!/usr/bin/env python
# -*- coding: utf-8 -*-
# ylgongPw @ 2020-02-18 14:04:55
from __future__ import unicode_literals
from __future__ import absolute_import

import threading
import time


class myThread(threading.Thread):
    def __init__(self,threadID,name,counter):
        super(myThread,self).__init__()
        self.threadID = threadID
        self.name = name
        self.counter = counter

    def run(self):
        print ("Starting {}".format(self.name))
        print_time(self.name, self.counter, 5)
        print ("Exiting {}".format(self.name))

def print_time(threadName, delay, counter):
    while counter:
        time.sleep(delay)
        print ("{}:{}".format(threadName,time.ctime(time.time())))
        counter -= 1

thread1 = myThread(1,"Thread-1", 1)
thread2 = myThread(2,"Thread-2", 2)

thread1.start()
thread2.start()

thread1.join()
thread2.join()





